window.onload = getTodoList();

const registerButton = document.getElementById(`register`);
registerButton.onclick = function() {
  const form = {};
  form.schedule = document.schedule_data.schedule.value;
  form.priority = document.schedule_data.priority.value;
  form.timeLimit = document.schedule_data.time_limit.value;
  const JsonData = JSON.stringify(form);
  fetch("/register", {
    method: "post",
    body: JsonData
  }).then(response => {
    if (response.status === 200) {
      alert("登録に成功しました");
      getTodoList();
    }
  });
};

function getTodoList() {
  fetch("/display")
    .then(response => {
      if (response.status === 200) {
        return response.json();
      }
    })
    .then(todoList => {
      const table = document.getElementById(`todo_list`);
      while (table.rows[0]) {
        table.deleteRow(0);
      }
      todoList.forEach(function(element) {
        let schedule = element.schedule;
        let priority = element.priority;
        let timeLimit = element.timeLimit;
        let id = element.id;
        let row = table.insertRow(-1);
        let cell = row.insertCell(-1);
        let text = document.createTextNode(schedule);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        text = document.createTextNode(priority);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        text = document.createTextNode(timeLimit);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        let button = document.createElement("button");
        button.type = "button";
        button.value = id;
        button.innerText = "削除";

        cell.appendChild(button);
        button.addEventListener("click", removeSchedule, false);
      });
    });
}

function removeSchedule(event) {
  const val = event.target.value;
  const form = { id: val };
  const JsonData = JSON.stringify(form);
  fetch("/remove", {
    method: "post",
    body: JsonData
  }).then(response => {
    if (response.status === 200) {
      alert("削除しました");
      getTodoList();
    }
  });
}
