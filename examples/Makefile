DOCTOOLS	=	$(HOME)/Projects/pandoctools
PANDOCFILTERS	=	$(HOME)/Projects/pandoctools/filters/codeblock/codeblock-filter

MDPARTS		=	notes.md geometry.md
MDPDFTARGETS	=	notes.md.pdf
MDHTMLTARGETS	=	notes.html
TEXPDFTARGETS	=
EXTRAARGS	=	$(foreach filter, $(PANDOCFILTERS), --filter $(filter))

include $(DOCTOOLS)/make/Makefile.in
