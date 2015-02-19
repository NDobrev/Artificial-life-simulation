function Point(x, y) {
    this.x = x;
    this.y = y;
}

Point.prototype.string = function () {
    return this.x + "_" + this.y;
}

Point.prototype.copy = function () {
    return new Point(this.x, this.y);
}

var xhr = new XMLHttpRequest();
xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) { // otgovora e gotov
        // tuka e json-a, moe6 da go prai6 kvot si iska6
        var matrix = JSON.parse(xhr.responseText); // tva e matricata
        
        var colors = matrix.Colors;
        for (var i = 0; i < colors.length; ++i) {
            for (var j = 0; j < colors[i].length; ++j) {
                var id = (new Point(i, j)).string();
                var element = document.getElementById(id);
                //console.log(element)
                element.style.backgroundColor = colors[i][j].toString();
            }
        }
    }
}
xhr.open('GET', 'http://localhost:8090/matrix', true);
xhr.send();

// fill table
function FillTable(width, height) {
    document.body = document.createElement("body");
    var table = document.body.appendChild(document.createElement("table"));
    for (var i = 0; i < width; ++i)
    {
        var tr = document.createElement("tr");
        table.appendChild(tr);
        for(var j=0; j<height; ++j)
        {
            var td = document.createElement("td");
            td.setAttribute("id", new Point(i, j).string());
            td.setAttribute("class", "free");
            tr.appendChild(td);
        }
    }
}

var interval = 500

function Fast(){
    interval = interval - 100
    clearInterval(myVar)
     rep()
}

function Slow(){
    interval = interval + 100
    clearInterval(myVar)
     rep()
}

function Reset(){
    xhr.open('POST', 'http://localhost:8090/reset', true);
    xhr.send();
}

function Do() {
    xhr.open('GET', 'http://localhost:8090/matrix', true);
    xhr.send();
    xhr.onreadystatechange()
}

function rep() {
        myVar = setInterval(Do, interval);
}
FillTable(140, 140)
rep()

//FillTable(70, 70)