document.getElementById("login").addEventListener("submit",async function(){
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value; 
    var data = await fetch("http://localhost:8080/api/login",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify({
            username:username,
            password:password
        })
    }).then(async (r)=>{
        if(!r.ok){
            if(r.status==401){
                Swal.fire({
                    title: "Username or Password Wrong",
                    text: "please try again",
                    icon: "error"
                  });
            }
        }
        else{
            await Swal.fire({
                title: "Success",
                text: "Good Jobs",
                icon: "success",
                confirmButtonText: 'OK',
              });
            window.location.href="index.html";
        }
    });
});