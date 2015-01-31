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
        alert(xhr.responseText);
    }
}
xhr.open('GET', 'http://localhost:8090/matrix', true);
xhr.send();

// fill table
function FillTable(mapSize) {
    document.body = document.createElement("body");
    var table = document.body.appendChild(document.createElement("table"));
    for (var i = 0; i < mapSize; ++i)
    {
        var tr = document.createElement("tr");
        table.appendChild(tr);
        for(var j=0; j<mapSize; ++j)
        {
            var td = document.createElement("td");
            td.setAttribute("id", new Point(i, j).string());
            td.setAttribute("class", "free");
            tr.appendChild(td);
        }
    }
}

//FillTable(70)