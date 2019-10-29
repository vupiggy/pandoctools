DOCTOOLS	=	$(HOME)/Projects/pandoctools
PANDOCFILTERS	=	$(DOCTOOLS)/filters/codeblock/codeblock-filter

MDPARTS		=	notes.md geometry.md
MDPDFTARGETS	=	notes.md.pdf
MDHTMLTARGETS	=	notes.html
TEXPDFTARGETS	=
EXTRAARGS	=	$(foreach filter, $(PANDOCFILTERS), --filter $(filter))

FIGURES_PDF=tangent_line.pdf inscribed_angle.pdf t4q25.pdf
FIGURES_SVG=$(patsubst %.pdf, %.svg, $(FIGURES_PDF))

include $(DOCTOOLS)/make/Makefile.in
