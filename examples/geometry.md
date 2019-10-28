# Code Snippet Example

```{.snippet}
{
    "path"    : "../filters/codeblock/figure/figure.go",
    "lang"    : "Go",
    "segment" : "figure"
}
```

# `amsthm` Example

## Circle

```{.theorem}
{
    "type" : "definition",
    "item" : "Plane",
    "text" : "In mathematics, a plane is a flat, two-dimensional surface that extends infinitely far."
}
```

```{.theorem}
{
    "type" : "definition",
    "item" : "Circle",
    "text" : "A circle is a shape consisting of all points in a plane that are a given distance from a given point, the centre; equivalently it is the curve traced out by a point that moves in a plane so that its distance from a given point is constant."
}
```

## Chord and tangent line

```{.theorem}
{
    "type" : "definition",
    "item" : "Chord",
    "text" : "A line segment whose endpoints lie on the circle, thus dividing a circle in two sements."
}
```

```{.theorem}
{
    "type" : "definition",
    "item" : "Tangent line",
    "text" : "A tangent line to a circle is a line that touches the circle at exactly one point, never entering the circle's interior."
}
```

# Figure Example

## Tangent Line

See `the figure below`{=html}`Figure~\ref{fig:tanl}`{=latex},
We have a theorem about tangent line to a circle:

```{.theorem}
{
    "type" : "theorem",
    "item" : "Tangent line to a circle",
    "text" : "A line is tangent to a circle, if and only if the line is perpendicular to the radius drawn to the point of tangency."
}
```

```{.figure}
{
    "path"    : "Figures/tangent_line",
    "caption" : "Tangent line to a circle",
    "label"   : "tanl",
    "options" : "scale=1",
    "place"   : "hb"
}
```

## Inscribed angle and central angle

```{.figure}
{
    "path"    : "Figures/inscribed_angle",
    "caption" : "Inscribed Angle and Central Angle",
    "label"   : "insc",
    "options" : "scale=1",
    "place"   : "hb"
}
```

## AMC8

With the code block below:

```Markdown
    ```{.figure}
    {
        "path"    : "Figures/t4q25",
        "caption" : "Test4 Q25",
        "label"   : "t4q25",
        "options" : "width=0.67\\textwidth",
        "place"   : "ht"
    }
    ```
```

You get `the figure below`{=html}`Figure~\ref{fig:t4q25}`{=latex}:

```{.figure}
{
    "path"    : "Figures/t4q25",
    "caption" : "Test4 Q25",
    "label"   : "t4q25",
    "options" : "width=0.67\\textwidth",
    "place"   : "ht"
}
```
