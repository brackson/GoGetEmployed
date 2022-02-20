(function($){
  $.fn.getFormData = function(){
    var data = {};
    var dataArray = $(this).serializeArray();
    for(var i=0;i<dataArray.length;i++){
      data[dataArray[i].name] = dataArray[i].value;
    }
    return data;
  }
})(jQuery);

$(document).ready(function () {
	// $("form").on("submit", function (e) {
		// e.preventDefault();
		// console.log(JSON.stringify($(this).getFormData()));
	// });

	// var myData = $("form$").getFormData();
});
