$(function() {                              
	$("p.play").click(function(event) {
	
		// Play
		if($(this).hasClass("play")){
			$(this).text("Pause");
			$(this).toggleClass("pause play");
			
			function reset(obj){
				obj.text("Play");
				obj.toggleClass("pause play");
			}
			
			obj = $(this);
			
			// Send the data using post
			var posting = $.post("play", {"name": $(this).parent().attr("id")}).done(
				function(data){
					setTimeout(reset(obj) , data.duration);
				}
			);
			
		// Pause
		}else if($(this).hasClass("pause")){
			$(this).text("Resume");
			$(this).toggleClass("pause resume");
			
			// Send the data using post
			var posting = $.post("pause", {"name": $(this).parent().attr("id")});
			
			
		// Resume
		}else if($(this).hasClass("resume")){
			$(this).text("Pause");
			$(this).toggleClass("pause resume");
			
			// Send the data using post
			var posting = $.post("resume", {"name": $(this).parent().attr("id")});
			
		}
	});			
			
		
			
	$("p.stop").click(function(event) {
		// Send the data using post
		var posting = $.post("stop", {"name": $(this).parent().attr("id")});
	});	
});
