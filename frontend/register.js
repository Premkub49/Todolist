document.getElementById("register").addEventListener("submit",async function(){
    event.preventDefault();
    var username = document.getElementById("username").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value; 
    var data = await fetch(`http://localhost:8080/api/register`,{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify({
            username:username,
            email:email,
            password:password
        })
    }).then(async (r)=>{
        if(!r.ok){
            Swal.fire({
                title: "Username Already Used",
                text: "please change username",
                icon: "error"
              });
        }
        else{
            await Swal.fire({
                title: "Success",
                text: "Good Jobs",
                icon: "success",
                confirmButtonText: 'OK',
              });
            window.location.href="login.html";
        }
    });
});