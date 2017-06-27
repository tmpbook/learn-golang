#!/bin/sh
############################################################
#
#version:0.2
#     set ip_forward = 1 and add version info
#version: 0.1 
#     Initial version,install docker-engine docker-compose docker-machine and kernel-4.1 for centos
#############################################################
set -e

version=0.2

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

echo_docker_as_nonroot() {
	if command_exists docker && [ -e /var/run/docker.sock ]; then
		(
			set -x
			$sh_c 'docker version'
		) || true
	fi
	your_user=your-user
	[ "$user" != 'root' ] && your_user="$user"
	# intentionally mixed spaces and tabs here -- tabs are stripped by "<<-EOF", spaces are kept in the output
	cat <<-EOF

	If you would like to use Docker as a non-root user, you should now consider
	adding your user to the "docker" group with something like:

	  sudo usermod -aG docker $your_user

	Remember that you will have to log out and back in for this to take effect!

	EOF
}

do_install() {
	case "$(uname -m)" in
		*64)
			;;
		*)
			cat >&2 <<-'EOF'
			Error: you are not using a 64bit platform.
			Docker currently only supports 64bit platforms.
			EOF
			exit 1
			;;
	esac

	if command_exists docker; then
		cat >&2 <<-'EOF'
			Warning: the "docker" command appears to already exist on this system.

			If you already have Docker installed, this script can cause trouble, which is
			why we're displaying this warning and provide the opportunity to cancel the
			installation.

			If you installed the current Docker package using this script and are using it
			again to update Docker, you can safely ignore this message.

			You may press Ctrl+C now to abort this script.
		EOF
		( set -x; sleep 20 )
	fi

	user="$(id -un 2>/dev/null || true)"

	sh_c='sh -c'
	if [ "$user" != 'root' ]; then
		if command_exists sudo; then
			sh_c='sudo -E sh -c'
		elif command_exists su; then
			sh_c='su -c'
		else
			cat >&2 <<-'EOF'
			Error: this installer needs the ability to run commands as root.
			We are unable to find either "sudo" or "su" available to make this happen.
			EOF
			exit 1
		fi
	fi

	curl=''
	if command_exists curl; then
		curl='curl -sSL'
	elif command_exists wget; then
		curl='wget -qO-'
	elif command_exists busybox && busybox --list-modules | grep -q wget; then
		curl='busybox wget -qO-'
	fi

	repo='main'

	# perform some very rudimentary platform detection
	lsb_dist=''
	dist_version=''
	if command_exists lsb_release; then
		lsb_dist="$(lsb_release -si)"
	fi
	if [ -z "$lsb_dist" ] && [ -r /etc/lsb-release ]; then
		lsb_dist="$(. /etc/lsb-release && echo "$DISTRIB_ID")"
	fi
	if [ -z "$lsb_dist" ]; then
		if [ -r /etc/centos-release ] || [ -r /etc/redhat-release ]; then
			lsb_dist='centos'
		fi
	fi
	if [ -z "$lsb_dist" ] && [ -r /etc/os-release ]; then
		lsb_dist="$(. /etc/os-release && echo "$ID")"
	fi

	lsb_dist="$(echo "$lsb_dist" | tr '[:upper:]' '[:lower:]')"

	case "$lsb_dist" in

		ubuntu)
			if command_exists lsb_release; then
				dist_version="$(lsb_release --codename | cut -f2)"
			fi
			if [ -z "$dist_version" ] && [ -r /etc/lsb-release ]; then
				dist_version="$(. /etc/lsb-release && echo "$DISTRIB_CODENAME")"
			fi
		;;

		centos)
			dist_version="$(rpm -q --whatprovides redhat-release --queryformat "%{VERSION}\n" | sed 's/\/.*//' | sed 's/\..*//' | sed 's/Server*//')"
		;;

	esac

	# install docker-compose
	install_docker_compose_ubuntu() {
		cat >&2 <<-'EOF'
		  Installing docker-compose in ubuntu
		EOF
		$sh_c 'sleep 3; apt-get install  -y -q --force-yes docker-compose' 
		$sh_c "chmod +x /usr/local/bin/docker-compose"
	}

	install_docker_compose_centos() {
		cat >&2 <<-'EOF'
		  Installing docker-compose in centos 
		EOF
		$sh_c 'sleep 3; yum -y -q --disableexcludes=all install  docker-compose' 
	}

	
	#install docker-machine
	install_docker_machine_ubuntu() {
		cat >&2 <<-'EOF'
		  Installing docker-machine in ubuntu
		EOF
		$sh_c 'sleep 3; apt-get install  -y -q  --force-yes docker-machine' 
	}

	install_docker_machine_centos() {
		cat >&2 <<-'EOF'
		  Installing docker-machine in centos 
		EOF
		$sh_c 'sleep 3; yum -y -q --disableexcludes=all install  docker-machine' 
	}


	# set nf_conntrack_max and and nf_conntrack hashsize
	set_nf_conntrack() {

		$sh_c "cat >/etc/modprobe.d/nf_conntrack.conf" <<-'EOF' 
		options nf_conntrack hashsize=131072
		EOF

                $sh_c 'sed -i "/net.netfilter.nf_conntrack_max/c net.netfilter.nf_conntrack_max = 104858886" /etc/sysctl.conf'
	}
	#set ip_forwarpv4.ip_forward
	set_ip_forward() {

                $sh_c 'sed -i "/net.ipv4.ip_forward/c net.ipv4.ip_forward = 1" /etc/sysctl.conf'
	}


	cat >&2 <<-'EOF'
	  Installing docker-engine
	EOF

	# Run setup for each distro accordingly
	case "$lsb_dist" in

		ubuntu)
			export DEBIAN_FRONTEND=noninteractive

			did_apt_get_update=
			apt_get_update() {
				if [ -z "$did_apt_get_update" ]; then
					( set -x; $sh_c 'sleep 3; apt-get update' )
					did_apt_get_update=1
				fi
			}

			# aufs is preferred over devicemapper; try to ensure the driver is available.
			if ! grep -q aufs /proc/filesystems && ! $sh_c 'modprobe aufs'; then
				if uname -r | grep -q -- '-generic' && dpkg -l 'linux-image-*-generic' | grep -q '^ii' 2>/dev/null; then
					kern_extras="linux-image-extra-$(uname -r) linux-image-extra-virtual"

					apt_get_update
					( set -x; $sh_c 'sleep 3; apt-get install -y -q '"$kern_extras" ) || true

					if ! grep -q aufs /proc/filesystems && ! $sh_c 'modprobe aufs'; then
						echo >&2 'Warning: tried to install '"$kern_extras"' (for AUFS)'
						echo >&2 ' but we still have no AUFS.  Docker may not work. Proceeding anyways!'
						( set -x; sleep 10 )
					fi
				else
					echo >&2 'Warning: current kernel is not supported by the linux-image-extra-virtual'
					echo >&2 ' package.  We have no AUFS support.  Consider installing the packages'
					echo >&2 ' linux-image-virtual kernel and linux-image-extra-virtual for AUFS support.'
					( set -x; sleep 10 )
				fi
			fi

			# install apparmor utils if they're missing and apparmor is enabled in the kernel
			# otherwise Docker will fail to start
			if [ "$(cat /sys/module/apparmor/parameters/enabled 2>/dev/null)" = 'Y' ]; then
				if command -v apparmor_parser >/dev/null 2>&1; then
					echo 'apparmor is enabled in the kernel and apparmor utils were already installed'
				else
					echo 'apparmor is enabled in the kernel, but apparmor_parser missing'
					apt_get_update
					( set -x; $sh_c 'sleep 3; apt-get install -y -q apparmor' )
				fi
			fi

			if [ ! -e /usr/lib/apt/methods/https ]; then
				apt_get_update
				( set -x; $sh_c 'sleep 3; apt-get install -y -q apt-transport-https ca-certificates' )
			fi
			if [ -z "$curl" ]; then
				apt_get_update
				( set -x; $sh_c 'sleep 3; apt-get install -y -q curl ca-certificates' )
				curl='curl -sSL'
			fi
			(
			set -x
			$sh_c "apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D"
			$sh_c "mkdir -p /etc/apt/sources.list.d"
			$sh_c "echo deb http://docker.mirrors.ucloud.cn/ubuntu/ ${lsb_dist}-${dist_version} ${repo} > /etc/apt/sources.list.d/ucloud-docker.list"
			$sh_c "echo deb [arch=amd64] http://docker.mirrors.ucloud.cn/ubuntu/ucloud ubuntu-ucloud main >> /etc/apt/sources.list.d/ucloud-docker.list"
			$sh_c 'sleep 3; apt-get update; apt-get install -y -q docker-engine'
			)
			install_docker_compose_ubuntu
			install_docker_machine_ubuntu
			echo_docker_as_nonroot
			set_nf_conntrack
			exit 0
			;;

		centos)
			$sh_c "cat >/etc/yum.repos.d/ucloud-docker-${repo}.repo" <<-EOF
			[docker-${repo}-repo]
			name=Docker ${repo} Repository
			baseurl=http://docker.mirrors.ucloud.cn/${lsb_dist}/\$releasever
			gpgcheck=0
			enabled=1
	     
			[UCloud-docker-${repo}-repo]
			name=UCloud Docker ${repo} Repository
			baseurl=http://docker.mirrors.ucloud.cn/${lsb_dist}/ucloud/\$releasever
			gpgcheck=0
			enabled=1
			EOF
			( set -x; $sh_c 'sleep 3; yum -y -q install docker-engine' )
			install_docker_machine_centos
			install_docker_compose_centos
			if [ "$dist_version" = 7 ]; then
				$sh_c 'yum -y update device-mapper'
				$sh_c 'systemctl enable docker.service'
			fi
			if [ "$dist_version" = 6 ]; then

				set_ip_forward 

				# centos6 open files ulimit defaults to 65536
				$sh_c 'sed -i "s/^other_args=\"/other_args=\"--default-ulimit nofile=65536:65536 /" /etc/sysconfig/docker'
				$sh_c 'service docker condrestart'
				cat >&2 <<-'EOF'
				  Installing kernel-4.1
				EOF
				( set -x; $sh_c 'sleep 3; yum -y -q --disableexcludes=all install  kernel-4.1.0' )
				if ! grubby --default-kernel | grep -q $(uname -r) ; then
					cat >&2 <<-'EOF'

					  ****************************
					  REBOOT system to take effect
					  ****************************
					EOF
				fi
			fi
			set_nf_conntrack
			exit 0
			;;
	esac
	# intentionally mixed spaces and tabs here -- tabs are stripped by "<<-'EOF'", spaces are kept in the output
	cat >&2 <<-'EOF'

	  Either your platform is not easily detectable, is not supported by this
	  installer script (yet - PRs welcome! [hack/install.sh]), or does not yet have
	  a package for Docker.  Please visit the following URL for more detailed
	  installation instructions:

	    https://docs.docker.com/en/latest/installation/

	EOF
	exit 1
}
if [ "$1" = "--version" -o "$1" = "-V" ]; then
echo "version: $version"
exit 0      
elif [ -n "$1" ]; then
echo "Description:"
echo "    This script attempts to install docker-engine docker-compose docker-machine"
echo "    and kernel-4.1.0"
echo "usage:"
echo "    $0 "
exit 0      
fi

# wrapped up in a function so that we have some protection against only getting
# half the file during "curl | sh"
do_install
