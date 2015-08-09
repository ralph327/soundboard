$(function() {
     $("form").each(function(){
          var form = $(this);
          
          var submit = form.find(":submit").attr("id");
          var action = form.attr("action");
          
          $("#"+submit).click(function(event) {
               // Stop form from submitting normally
               event.preventDefault();
               
               var obj = {};
                               
               form.find(":input").not('textarea,[type="checkbox"],[type="submit"]').each(function(){
                    var input = $(this);
                    obj[input.attr("id")] = input.val();
               });
               
               // CHECKBOXES
               var checkboxes = {};
               var names = new Array();
               // Initiate 
               form.find('[type="checkbox"]:checked').each(function(){
                    var input = $(this);
                    var name = input.attr("name");
                    
                    if($.inArray(name, checkboxes) == -1){
                         checkboxes[name] = new Array();
                         names.push(name);
                    }
               });
               
               // Fill checkboxes array
               form.find('[type="checkbox"]:checked').each(function(){
                    var input = $(this);
                    var name = input.attr("name");

                    checkboxes[name].push(input.val());
               });
               
               // Add to obj
               for (index = 0; index < names.length; ++index) {
                    var n = names[index];
                    obj[n] = checkboxes[n];
               }

               // TEXTAREA             
               form.find("textarea").each(function(){
                    var input = $(this);
                    obj[input.attr("id")] = input.val();
               });
                             
               var json = JSON.stringify(obj);   
               
               alert(json);
               
               // Send the data using post
               var posting = $.post(action, json).done(function( data ) {
                    if(data.error == "none"){
                         $(':input', form)
                              .not(':button, :submit, :reset, :hidden')
                              .val('')
                              .removeAttr('checked')
                              .removeAttr('selected');
                         $('textarea', form).val('');    
                         flash($(".flash-success"), data.success);
                    }else{
                         flash($(".flash-error"), data.error);
                    }
               });
          });
     });
});
