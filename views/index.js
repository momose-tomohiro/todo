const todos = new Vue({
  el: "#todoPage",
  vuetify: new Vuetify(),
  data: {
      header:[
        {
          text:'予定',
          align: 'center',
          value: 'schedule'
        },
        {
          text:'優先度',
          align: 'center',
          value:'priority'
        },
        {
          text:'期限',
          align: 'center',
          value:'time_limit'
        },
        {
          text:'削除',
          align: 'center',
          value:'action',
          sortable:false
        }
      ],
      todoList: [
        {
          id: '',
          schedule: '',
          priority: '',
          time_limit: ''
        },
      ],
      removeConfirmationDialog:false,
      finishedDialog:false,
      finishType:'',
      errorDialog:false,
      removeID:''
    },
  mounted: function () {
    this.display();
  },
  methods: {
    display: function () {
      fetch("/todos")
        .then((response) => {
          if (response.ok) {
            return response.json();
          }
        })
        .then((result) => {
          todos.todoList = result;
        })
        .catch((err) => {
          this.errorDialog = true;
          console.log(err);
        });
    },
    register: function () {
      const form = {
        schedule: this.todoList.schedule,
        priority: this.todoList.priority,
        time_limit: this.todoList.time_limit,
      };

      fetch("/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(form),
      })
        .then((response) => {
          if (response.ok) {
            this.finishType = '登録';
            this.finishedDialog = true;
            this.todoList.schedule = '';
            this.todoList.time_limit = '';
            this.display();
          }
        })
        .catch((err) => {
          this.errorDialog = true;
          console.log(err);
        });
    },
    remove: function () {
      fetch(`/todos?id=${this.removeID}`, {
        method: "DELETE",
      })
        .then((response) => {
          if (response.ok) {
            this.removeConfirmationDialog = false;
            this.finishType = '削除';
            this.finishedDialog = true;
            this.display();
          }
        })
        .catch((err) => {
          this.errorDialog = true;
          console.log(err);
        });
    },
  },
});
