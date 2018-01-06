const input = document.getElementById("searchInput");
const table = document.getElementById("tableList");

function search() {
    // var input, filter, ul, li, a, i;
    // input = document.getElementById("myInput");
    // filter = input.value.toUpperCase();
    // ul = document.getElementById("myUL");
    // li = ul.getElementsByTagName("li");
    // for (i = 0; i < li.length; i++) {
    //     a = li[i].getElementsByTagName("a")[0];
    //     if (a.innerHTML.toUpperCase().indexOf(filter) > -1) {
    //         li[i].style.display = "";
    //     } else {
    //         li[i].style.display = "none";

    //     }
    // }
}

function openSpec(section) {
	window.open(
		"https://github.com/KhronosGroup/glTF/blob/master/specification/2.0/README.md#" + section,
		"_blank");
}

function init() {
	document.getElementById("lastUpdated").innerHTML = LAST_UPDATE;
	input.focus();

	// adds items to table
	for (var i = 0; i < ENUMS.length; i++) {
		var row = table.insertRow(-1);
		row.className = "tableRow";
		(row.insertCell(0)).innerHTML = ENUMS[i].value;
		(row.insertCell(1)).innerHTML = ENUMS[i].name;
		var typeCell = row.insertCell(2);
		typeCell.innerHTML = "";
		for (var j = 0; j < ENUMS[i].types.length; j++) {
			typeCell.innerHTML += "<a href=\"#\" onclick=\"openSpec('" + ENUMS[i].types[j].link +
			"')\">" + ENUMS[i].types[j].name + "</a> [<span>"+ ENUMS[i].types[j].type +
			"</span>] (<text>" + ENUMS[i].types[j].required + "</text>)"

			if (ENUMS[i].types.length != j + 1) {
				typeCell.innerHTML += "<br>";
			}
		}
	}
}