// BUG:
// Currently being imported in index.html, but not being used
// 

console.log('Test');

$(".resume-button").click(function(){
    $.ajax({url: "http://localhost:1323/resume_report", success: function(result){
        console.log('Resuming Report: ', result);
    }});
});

$(".stop-button").click(function(){
    $.ajax({url: "http://localhost:1323/halt_report", success: function(result){
        console.log('Halting Report: ', result);
    }});
});

