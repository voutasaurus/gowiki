var ToC =
  "<nav role='navigation' class='table-of-contents'>" +
    "<h5>Contents</h5>" +
    "<ol>";
	
var el, title, link;

$("article h3").each(function() {

	el = $(this);
	title = el.text();
	link = "#" + el.attr("id");

	newLine =
	"<li>" +
	  "<a href='" + link + "'>" +
		title +
	  "</a>" +
	"</li>";

	ToC += newLine;
  
});

ToC +=
   "</ol>" +
  "</nav>";
  
$("article").prepend(ToC);