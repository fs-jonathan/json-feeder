window.onload = function (e) {
  liff.init(function (data) {
    initializeApp(data);
  }, function() {
    $('#error').append("Liff Error");
    // getData();
  });
};

function initializeApp(data) {
  if (data) {
    $.ajax({
      url: '/loginLiff',
      type: 'POST',
      dataType: 'json',
      data: {
        "lineUserId": data.context.userId
      },
      timeout: 5000
    }).done(function(data) {
      getData();
    }).fail(function(XMLHttpRequest, textStatus, errorThrown) {
      showLogin();
    })
  } else {
    $('#error').append("Wrong Page Access");
  }
}

function showLogin() {
  $.ajax('login.html', {
          timeout : 1000, // 1000 ms
          datatype:'html'
      }).done(function(data) {
        $('#login').append(data);
      });
}

function getData() {
  window.location.href = "getRecord.html";
}
