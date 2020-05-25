window.onload = getTodoList();
const table = document.getElementById(`todo_list`);

const registerButton = document.getElementById(`register`);
registerButton.onclick = function() {
  const form = {
    schedule : document.schedule_data.schedule.value,
    priority : document.schedule_data.priority.value,
    timeLimit : document.schedule_data.time_limit.value
  };
  
  fetch('/todos', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(form)
  }).then(response => {
    if (response.ok) {
      alert('登録に成功しました');
      document.schedule_data.schedule.value = "";
      document.schedule_data.time_limit.value = "";
      getTodoList();
    }
  }).catch((err) => {
    console.log(err);
  })
};

function getTodoList() {
  fetch('/todos')
    .then(response => {
      if (response.ok) {
        return response.json();
      }
    })
    .then(todoList => {
      table.innerHTML = ""
      for(todo of todoList){
        const tr = document.createElement("tr")
        const schedule = document.createElement("td")
        schedule.textContent = todo.schedule
        tr.appendChild(schedule)
        const priority = document.createElement("td")
        priority.textContent = todo.priority
        tr.appendChild(priority)
        const timeLimit = document.createElement("td")
        timeLimit.textContent = todo.timeLimit
        tr.appendChild(timeLimit)
        const button = document.createElement('button');
        button.type = 'button';
        button.value = todo.id;
        button.textContent = '削除';
        button.addEventListener('click', removeSchedule, false);
        tr.appendChild(button)

        table.appendChild(tr)
      }
      /*以前
      todoList.forEach(function(element) {
        let row = table.insertRow(-1);
        let cell = row.insertCell(-1);
        let text = document.createTextNode(element.schedule);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        text = document.createTextNode(element.priority);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        text = document.createTextNode(element.timeLimit);
        cell.appendChild(text);
        cell = row.insertCell(-1);
        let button = document.createElement('button');
        button.type = 'button';
        button.value = element.id;
        button.innerText = '削除';

        cell.appendChild(button);
        button.addEventListener('click', removeSchedule, false);
      });
      */
    }).catch((err) => {
      console.log(err);
    })
}

function removeSchedule(event) {
  const id = event.target.value;
  fetch(`/todos?id=${id}`, {
    method: 'DELETE'
  }).then(response => {
    if (response.ok) {
      alert('削除しました');
      getTodoList();
    }
  }).catch((err) => {
    console.log(err);
  })
}
