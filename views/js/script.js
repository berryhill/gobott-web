console.log('Test');

var seconds = 1;
var url = "http://localhost:1323"

$(".resume-button").click(function(){
    $.ajax({url: url + "/resume_report", success: function(result){
        console.log('Resuming Report: ', result);
    }});
});

$(".stop-button").click(function(){
    $.ajax({url: url + "/halt_report", success: function(result){
        console.log('Halting Report: ', result);
    }});
})

$(".set-timer").click(function(){
    $.ajax({url: url + "/set_timer/" + '#seconds', success: function(result){
        console.log('Setting Timer: ', result);
    }});
});