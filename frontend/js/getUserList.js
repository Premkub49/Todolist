function successTask(id) {
  fetch(`${window.backend}/api/deleteTask`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: id,
    }),
  }).then((r) => {
    if (r.ok) {
      document.getElementById(`detailCard-${id}`).remove();
      window.getUserList();
    }
  });
}
function editTask(id) {
  let listname = document.getElementById("edit-namelist").value;
  let deadline = document.getElementById("edit-deadline").value;
  let detail = document.getElementById("edit-detail").value;
  fetch(`${window.backend}/api/updateTask`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id: id,
      listname: listname,
      deadline: deadline,
      detail: detail,
    }),
  }).then((r) => {
    if (r.ok) {
      document.getElementById("editTask").remove();
      window.getUserList();
    }
  });
}
function editTaskBox(id, listname, deadline, detail) {
  let div = document.createElement("div");
  div.id = `editTask`;
  div.className = "Card-List";
  div.style.zIndex = "9";
  div.style.visibility = "visible";
  deadline = deadline.replace("Z", "");
  div.innerHTML = `
            <ion-icon onclick = "document.getElementById('editTask').remove();" name="close-outline"
                style="grid-area: 2/2/3/3; width: 25px; height: 25px; cursor: pointer"></ion-icon>
            <label
                style="grid-area: 2/1/4/3; display: flex; justify-content: center; align-items: center; font-size: x-large; font-weight: bolder;">Edit List</label>
            <form id="editListForm" action="#" onsubmit="editTask(${id});"
                style="grid-area: 4/1/8/3; display: grid; grid-template-rows: 20% 20% 45% 15%; margin-left:5%;">
                <div style="grid-area: 1/1/2/2" class="input-div">
                    <p class="input-label">Name</p>
                    <input class="input-box" type="text" name="namelist"
                        id="edit-namelist" value = "${listname}" required />
                </div>
                <div style="grid-area: 2/1/3/2" class="input-div">
                    <p class="input-label">Deadline</p>
                    <input type="datetime-local" name="deadline"
                        id="edit-deadline" class="input-box" value= "${deadline}" />
                </div>
                <div class="input-div"
                    style="grid-area:3/1/4/2; max-width: 90%; width:100%; height: 100%; display:flex; flex-direction: column;">
                    <p class="input-label">Deatail</p>
                    <input type="text" name="detail" id="edit-detail"
                        class="input-box" style="width: 90%; height: 90%;" value="${detail}" />
                </div>
                <button class="btn btn-success"
                    style="grid-area:4/1/5/2; width: 80px; height: 45px; margin:auto;">EditList</button>
            </form>`;
  document.getElementById("bodyList").appendChild(div);
  document.getElementById(`detailCard-${id}`).style.visibility = "hidden";
}
window.getUserList = async () => {
  await fetch(`${window.backend}/api/getUserTask`, {
    method: "POST",
    headers: { "Content-Type": "appilcation/json" },
    body: JSON.stringify({
      username: window.username,
    }),
  }).then(async (r) => {
    r = await r.json();
    document.getElementById("table-body").innerHTML = "";
    for (var i = 1; i <= r.length; i++) {
      let data = r[i - 1];
      data.deadline = data.deadline.replace("T", " ");
      data.deadline = data.deadline.replace(":00Z", "");
      var swapdeadline = [8, 9, 7, 5, 6, 4, 0, 1, 2, 3, 10, 11, 12, 13, 14, 15];
      var copydeadline = data.deadline;
      var newdeadline = "";
      for (var j = 0; j < 16; j++) {
        newdeadline += copydeadline[swapdeadline[j]];
      }
      data.deadline = newdeadline;
      let tr = document.createElement("tr");
      tr.innerHTML += `<td class="bodies">${data.listname}</td>`;
      tr.innerHTML += `<td class="bodies">${data.deadline}</td>`;
      tr.innerHTML += `<td class="bodies" onclick="document.getElementById('detailCard-${data.id}').style.visibility = 'visible';" style="cursor:pointer; text-align:center;" onmouseover="this.style.textDecoration='underline';" onmouseout="this.style.textDecoration='none';">Detail</td>`;
      document.getElementById("table-body").appendChild(tr);
      let div = document.createElement("div");
      div.id = `detailCard-${data.id}`;
      div.className = "Card-List";
      div.style.gridTemplateRows = "10% 15% 15% 40% 15% 5%";
      div.style.gridTemplateColumns = "50% 40% 8% 2%";
      div.innerHTML += `
            <ion-icon name="close-outline"
                style="grid-area: 1/3/2/4; width: 25px; height: 25px; cursor: pointer; margin:auto;" onclick="document.getElementById('detailCard-${data.id}').style.visibility = 'hidden';"></ion-icon>
            <label style="grid-area: 2/1/3/4; display: flex; justify-content: center; align-items: center; font-size: xx-large; font-weight: bolder;">${data.listname}</label>
                <div style="grid-area: 3/1/4/4; margin-left:20px;"  class="input-div">
                    <p class="input-label">Deadline</p>
                    <p>${data.deadline}</p>
                </div>
                <div class="input-div" style="grid-area:4/1/5/4; margin-left:20px; max-width: 90%; width:100%; height: 100%; display:flex; flex-direction: column;" >
                    <p class="input-label">Detail</p>
                    <p>${data.detail}</p>
                </div>
                <button class="btn btn-success" style="grid-area:5/1/6/2; width: 80px; height: 45px; margin:auto;" onclick="successTask(${data.id})">Success</button>
                <button class="btn btn-danger" style="grid-area:5/2/6/4; width: 80px; height: 45px; margin:auto;" onclick='editTaskBox(${data.id},"${data.listname}","${data.deadline}","${data.detail}")'>EditList</button>`;
      document.getElementById("bodyList").appendChild(div);
    }
  });
};
