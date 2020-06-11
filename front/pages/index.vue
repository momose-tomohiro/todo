<template>
  <v-app>
    <v-container>
      <v-row justify="center" align-content="center">
        <v-col cols="10">
          <h1 class="text-center">todo</h1>
          <v-data-table :headers="header" :items="todoList" item-key="id">
            <template v-slot:item.action="{ item }">
              <v-btn
                small
                color="error"
                @click="(removeConfirmationDialog = true), (removeID = item.id)"
                >削除</v-btn
              >
            </template>
          </v-data-table>

          <form id="schedule_data">
            <div class="schedule">
              <v-text-field
                :counter="255"
                name="schedule"
                v-model="todoList.schedule"
                label="予定を入力してください"
              />
            </div>
            優先度
            <v-radio-group v-model="todoList.priority">
                <v-radio label="高" value="高"></v-radio>
                <v-radio label="中" value="中"></v-radio>
                <v-radio label="低" value="低"></v-radio>
            </v-radio-group>
            <div class="time-limit">
              <v-menu>
                <template v-slot:activator="{on}">
                  <v-btn icon color="primary" v-on="on">
                    <v-icon>mdi-calendar</v-icon>
                  </v-btn> 
                </template>
                <v-date-picker  locale="ja" v-model="todoList.time_limit"/>
              </v-menu>
              
              <v-text-field
                :counter="255"
                name="time_limit"
                v-model="todoList.time_limit"
                label="期限を入力してください"
              />
            </div>
            <div class="text-center">
              <v-btn color="primary" id="register" v-on:click="register">
                登録
              </v-btn>
            </div>
          </form>
        </v-col>
      </v-row>

    <v-dialog v-model="removeConfirmationDialog" persistent max-width="350">
        <v-card>
          <v-row justify="center" align-content="center" no-gutters>
            <v-card-title>削除確認</v-card-title>
            <v-card-text class="text-center">本当に削除しますか？</v-card-text>
            <v-card-actions>
              <v-btn @click="removeConfirmationDialog = false">キャンセル</v-btn>
              <v-btn color="error" @click="remove()">削除</v-btn>
            </v-card-actions>
          </v-row>
        </v-card>
      </v-dialog>

      <v-dialog v-model="finishedDialog" persistent max-width="350">
        <v-card>
          <v-row justify="center" align-content="center" no-gutters>
            <v-card-title></v-card-title>
            <v-card-text class="text-center">
              <v-icon color="green">mdi-checkbox-marked-circle</v-icon>
              {{finishType}}しました。
            </v-card-text>
            <v-card-actions>
              <v-btn @click="finishedDialog = false">閉じる</v-btn>
            </v-card-actions>
          </v-row>
        </v-card>
      </v-dialog>

      <v-dialog v-model="errorDialog" persistent max-width="350">
        <v-card>
          <v-row justify="center" align-content="center" no-gutters>
            <v-card-title></v-card-title>
            <v-card-text class="text-center">
              <v-icon color="error">mdi-cancel</v-icon>
              エラーが発生しました。
            </v-card-text>
            <v-card-actions>
              <v-btn @click="errorDialog = false">閉じる</v-btn>
            </v-card-actions>
          </v-row>
        </v-card>
      </v-dialog>
      
    </v-container>
  </v-app>
</template>

<script>
import axios from 'axios';

export default {

  async asyncData({$axios}){
    const data = await $axios.$get('/todos')
    .catch((err) => {
          console.log(err);
        });
    return{todoList:data}
  },
  data() {
    return {
      header: [
        {
          text: "予定",
          align: "center",
          value: "schedule"
        },
        {
          text: "優先度",
          align: "center",
          value: "priority"
        },
        {
          text: "期限",
          align: "center",
          value: "time_limit"
        },
        {
          text: "削除",
          align: "center",
          value: "action",
          sortable: false
        }
      ],
      todoList: [
        {
          id: "",
          schedule: "",
          priority: "",
          time_limit: ""
        }
      ],
      removeConfirmationDialog: false,
      finishedDialog: false,
      finishType: "",
      errorDialog: false,
      removeID: ""
    };
  },

  methods:{
    display:function(){
      axios.get('/todos')
      .then((response) => {
          if (response.statusText === 'OK') {
            this.todoList = response.data
          }
        })
        .catch((err) => {
          this.errorDialog = true;
          console.log(err);
        });
    },
    register:function(){
      axios.post('/todos',{
        schedule: this.todoList.schedule,
        priority: this.todoList.priority,
        time_limit: this.todoList.time_limit
      }
      )
      .then((response) => {
          if (response.statusText === 'OK') {
            this.finishType = '登録'
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
    remove:function(){
      this.removeConfirmationDialog = false;
      axios.delete('/todos/' + this.removeID)
      .then((response) => {
          if (response.statusText === 'OK') {
            this.finishType = '削除';
            this.finishedDialog = true;
            this.display();
          }
        })
        .catch((err) => {
          this.errorDialog = true;
          console.log(err);
        });
    }
  }
};
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
