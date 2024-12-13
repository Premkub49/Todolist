document.getElementById("addListForm").addEventListener("submit",function(){
    event.preventDefault();
    let listname = document.getElementById("namelist").value;
    let deadline = document.getElementById("deadline").value;
    let detail = document.getElementById("detail").value;
    deadline = deadline.replace("T"," ")
    deadline+=":00"
    fetch(`${window.backend}/api/createTask`,{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify({
            "listname":listname,
            "deadline":deadline,
            "detail":detail,
            "username":window.username
        })
    }).then(async (r)=>{
        if(!r.ok){
            //console.log(r)
        }
        else{
            window.getUserList();
            document.getElementById("addCard").style.visibility = "hidden";
        }
    })
})