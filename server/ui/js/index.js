$(document).ready(function() {
    var text = GetQueryString("text");
    if (text != null) {
        $('#text').val(text);
        Parse(text);
    }
    $('#submit').click(function() {
        var text = $('#text').val().trim();
        if (text != "") {
            Parse(text);
        }
    });
    $("#text").on('keypress', function(e) {
        if (e.keyCode != 13) return;
        var text = $('#text').val().trim();
        if (text != "") {
            Parse(text);
        }
    });
});

function Parse(text) {
    $('#editor_holder').html("<h4>loading...</h4>");
    $("#visual").html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/?text="+encodeURIComponent(text), cache: false,
        success: function(result) {
            $('#editor_holder').jsonview(result);
            visual(result);
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}

function one(k, v) {
    return "<fieldset><legend class=\"label label-info left\">"+
        k+"</legend>"+v+"</fieldset>";
}

function visual(doc) {
    var html = "";
    for (var i=0;i<doc.length;i++) {
        html += one("i", i);
        for (k in doc[i]) {
            html += one(k, doc[i][k]);
        }
    }
    $("#visual").html(html);
}
