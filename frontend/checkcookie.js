function getCookie(name){
    let nameEq = name + "=";
    let ca = document.cookie.split(';');
    for(var i=0;i<ca.length;i++){
        var c = ca[i].trim();
        if(c.indexOf(nameEq)===0)return c.substring(nameEq.length,c.length);
    }
    return null;
}
let token = getCookie("token");
if(token===null){
    window.location.href = "login.html";
}