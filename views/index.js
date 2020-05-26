const todos = new Vue({
  el:'#todoPage',
  data:{
    todoList:""
  },
  methods:{
    display:onload=function(){
      fetch('/todos')
        .then(response => {
          if (response.ok) {
            return response.json();
          }
        })
        .then(result => {
          todos.todoList = result
        }).catch((err) => {
          console.log(err);
        })
    },
    register:function(){
      const form = {
        schedule : this.schedule,
        priority : this.priority,
        timeLimit : this.time_limit
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
          this.schedule = "";
          this.time_limit = "";
          this.display()
        }
      }).catch((err) => {
        console.log(err);
      })
    },
    remove:function(event){
      const id = event.target.value;
      fetch(`/todos?id=${id}`, {
        method: 'DELETE'
      }).then(response => {
        if (response.ok) {
          alert('削除しました');
          this.display()
        }
      }).catch((err) => {
        console.log(err);
      })
    }
  }
})
