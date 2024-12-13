function successTask(id){
    fetch(`${window.backend}/api/deleteTask`,{
        "method":"DELETE",
        "headers":{
            "Content-Type":"application/json"
        },
        "body":JSON.stringify({
            "id" : id
        })
    }).then((r)=>{
        if(r.ok){
            document.getElementById(`detailCard-${id}`).remove()
            window.getUserList();
        }
    })
}

window.getUserList = async()=>{
    await fetch(`${window.backend}/api/getUserTask`,{
        method:"POST",
        headers:{"Content-Type":"appilcation/json"},
        body:JSON.stringify({
            "username":window.username
        })
    }).then(async(r)=>{
        r = await r.json();
        document.getElementById("table-body").innerHTML = "";
        for(var i=1;i<=r.length;i++){
            let data = r[i-1];
            let tr = document.createElement("tr");
            tr.innerHTML+=`<td class="bodies">${data.listname}</td>`
            tr.innerHTML+=`<td class="bodies">${data.deadline}</td>`
            tr.innerHTML+=`<td class="btn btn-secondary bodies" onclick="document.getElementById('detailCard-${data.id}').style.visibility = 'visible';" style="cursor:pointer; width:80px; text-align:center;">Detail</td>`
            document.getElementById("table-body").appendChild(tr);
            let div = document.createElement("div");
            div.id = `detailCard-${data.id}`;
            div.className = "Card-List";
            div.style.gridTemplateRows = "2% 15% 15% 47% 15% 5%";
            div.style.gridTemplateColumns = "50% 45% 3% 2%";
            div.innerHTML += `
            <ion-icon name="close-outline"
                style="grid-area: 2/3/3/4; width: 25px; height: 25px; cursor: pointer" onclick="document.getElementById('detailCard-${data.id}').style.visibility = 'hidden';"></ion-icon>
            <label style="grid-area: 2/1/3/4; display: flex; justify-content: center; align-items: center; font-size: xx-large; font-weight: bolder;">${data.listname}</label>
                <div style="grid-area: 3/1/4/4"  class="input-div">
                    <p class="input-label">Deadline</p>
                    <p>${data.deadline}</p>
                </div>
                <div class="input-div" style="grid-area:4/1/5/4; max-width: 90%; width:100%; height: 100%; display:flex; flex-direction: column;" >
                    <p class="input-label">Detail</p>
                    <p>${data.detail}</p>
                </div>
                <button class="btn btn-success" style="grid-area:5/1/6/2; width: 80px; height: 45px; margin:auto;" onclick="successTask(${data.id})">Success</button>
                <button class="btn btn-danger" style="grid-area:5/2/6/4; width: 80px; height: 45px; margin:auto;">EditList</button>`;
            document.getElementById("bodyList").appendChild(div);
        }
    })
}