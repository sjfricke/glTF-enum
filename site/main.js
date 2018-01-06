const enumList = [
	{"value" : 5121, "name" : "UNSIGNED_BYTE", "types" : [
		{ "name" : "accessor.componentType", "required" : "yes", "type" : "integer"},
		{ "name" : "indices.componentType", "required" : "yes", "type" : "integer"}]},
	{"value" : 34962, "name" : "ARRAY_BUFFER", "types" : [
		{ "name" : "bufferView.target", "required" : "yes", "type" : "integer"}]},
	{"value" : 5, "name" : "TRIANGLE_STRIP", "types" : [
		{ "name" : "primitive.mode", "required" : "no", "type" : "integer", "default" : "TRIANGLE"}]}
];


const input = document.getElementById("searchInput");

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

function init() {
	input.focus();
	// var list = document.getElementById("myUL");
	// for (var i = 0; i < gitDictionary.length; i++) {
	// 	list.innerHTML += "<li><a href='#'><span style='color:red'>"
	// 					   + gitDictionary[i].command
	// 					   + "</span> - " + gitDictionary[i].description +"</a></li>";
	// }

	// https://github.com/KhronosGroup/glTF/blob/master/specification/2.0/README.md#bufferviewtarget
}