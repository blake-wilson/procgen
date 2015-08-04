// initialization
var canvas = document.getElementById("map_canvas");
var ctx = canvas.getContext("2d");
ctx.beginPath();
ctx.moveTo(0,0);

canvasHeight = document.getElementById("map_canvas").height;

function drawLine(start, end) {
	ctx.beginPath();
	ctx.moveTo(start.x, canvasHeight);
	ctx.lineTo(start.x, start.y);
	ctx.lineTo(end.x, end.y);
	ctx.lineTo(end.x, canvasHeight);
	ctx.fillStyle = start.color;
	ctx.closePath();
	ctx.fill();
	console.log(start.color);
}

function getValues(text) {
	text = document.getElementById("text").value;
	// assume format x,y\n for each coordinate
	coordinates = text.split("\n");
	tileAttrs = [];
	for (i = 0; i < coordinates.length; i++) {
		splitAttrs = coordinates[i].split(",");
		console.log(splitAttrs)
		if (splitAttrs.length < 3) {
			continue
		}
		var attrs = {
			x: parseInt(splitAttrs[0]),
			y: parseInt(splitAttrs[1]),
			color: '#' + splitAttrs[2].trim(),
		};
		tileAttrs.push(attrs);
	}
	return tileAttrs
}

function doIt() {
	ctx.clearRect(0, 0, canvas.width, canvas.height);
	values = [];
	values = getValues();
	for (i = 0; i < values.length - 1; i++) {
		var fst = {
			x: values[i].x,
			y: values[i].y,
			color: values[i].color,
		};
		var snd = {
			x: values[i + 1].x,
			y: values[i + 1].y,
			color: values[i + 1].color,
		};
		drawLine(fst, snd);
	}
}
