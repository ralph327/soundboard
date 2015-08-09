$(function() {
     $("li").each(function(){
          var block = $(this);
                    
          var name = block.attr("id");
          
          block.click(function(event) {
			// Send the data using post
               var posting = $.post("play", {"name": name});
		});
	});
});
