procgen

procgen is a tool to create 2-dimensional procedurally generated
terrain.  Procgen provides adjustable parameters for attributes
such as steepness and texture.


Functions require the context of each tile generated up to the point of the
current tile.

Each cell attribute is calcuated this way.  The value for an attribute is
therefore determined by a function f and the pervious tile values x1, x2, ...,
xn.
This can be represented in function terms as an = f(x1, x2, x3, ..., xn-1)
where an is the nth attribute.

This means that the value of an attribute can be a function of all attributes of
all preceding values.


Example of a function

f(x1, x2, x3, ...,xn) = x1.temp + x2.temp + ... + xn.temp

here, the function f determines the attribute temp for a given cell, and it does
so by summing the temp attributes of all preceding tiles
