DOCTOOLS	=	..
PANDOCFILTERS	=	$(DOCTOOLS)/filters/codeblock/codeblock-filter $(DOCTOOLS)/filters/image/image-filter

MDPARTS		=	notes.md geometry.md
MDPDFTARGETS	=	notes.md.pdf
MDHTMLTARGETS	=	notes.html
TEXPDFTARGETS	=
EXTRAARGS	=	$(foreach filter, $(PANDOCFILTERS), --filter $(filter)) --filter pandoc-crossref
RESOURCEPATH	=	"--resource-path=."

FIGURES_PDF=tangent_line.pdf inscribed_angle.pdf t4q25.pdf
FIGURES_SVG=$(patsubst %.pdf, %.svg, $(FIGURES_PDF))

include $(DOCTOOLS)/make/Makefile.in
