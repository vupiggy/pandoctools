FROM ubuntu:eoan
MAINTAINER Luke Huang <lukehuang.ca@me.com>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -yq update &&                                               \
    apt-get install -qy --no-install-recommends --fix-missing           \
        texlive                                                         \
        texlive-base                                                    \
        texlive-binaries                                                \
        texlive-lang-chinese                                            \
        texlive-latex-base                                              \
        texlive-latex-extra                                             \
        texlive-science                                                 \
        texlive-xetex                                                   \
        texlive-bibtex-extra                                            \
        pandoc                                                          \
        node-katex                                                      \
        make &&                                                         \
    apt-get -yq autoremove &&                                           \
    apt-get clean -y &&                                                 \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*