document.getElementById("logout").addEventListener("click",function(){
    document.cookie = "token= ;path=/; expires=Thu, 18 Dec 2013 12:00:00 UTC;";
    window.location.href="login.html";
})