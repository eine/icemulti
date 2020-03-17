docker_guiapp.sh --rm -itv /$(pwd)://src fedora bash -c "\
  dnf install -y redhat-lsb.i686 && \
  ./iCEcube2setup_Sep_12_2017_1708"
