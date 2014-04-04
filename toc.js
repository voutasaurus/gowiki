var ToC =
  "<nav role='navigation' class='table-of-contents'>" +
    "<h2>On this page:</h2>" +
    "<ul>";
	
var el, title, link;

$("article h3").each(function() {

  el = $(this);
  title = el.text();
  link = "#" + el.attr("id");

});