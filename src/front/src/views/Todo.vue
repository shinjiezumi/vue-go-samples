<template>
  <v-content>
    <v-row class="text-center" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <h1>TodoList</h1>
      </v-col>
    </v-row>
    <v-row align="center" justify="center">
      <v-col class="mb-3" cols="12" sm="8">
        <v-expansion-panels multiple>
          <v-expansion-panel v-for="todo in this.getTodoList" :key="todo.id">
            <v-expansion-panel-header :disable-icon-rotate="!!todo.finished_at">
              <div class="d-flex todo--header">
                <div class="mb-3">
                  <span class="mr-3"><i class="fas fa-stopwatch mr-3"></i>{{formatDate(todo.limit_date)}}</span>
                </div>
                <h2>{{todo.title}}</h2>
              </div>
              <template v-slot:actions v-if="todo.finished_at">
                <v-icon color="teal">mdi-check</v-icon>
              </template>
            </v-expansion-panel-header>
            <v-expansion-panel-content class="todo--memo">
              {{todo.memo}}
              <div class="text-right">
                <span class="mr-3">
                  <v-btn v-if="!todo.finished_at" class="mx-2" fab dark @click="finish(todo.id)" color="primary"
                         x-small>
                    <v-icon dark>mdi-check</v-icon>
                  </v-btn>
                </span>
                <span class="mr-3">
                  <TodoForm
                      :todo-id="todo.id"
                      :todo-title="todo.title"
                      :todo-memo="todo.memo"
                      :todo-limit-date="todo.limit_date"
                  />
                </span>
                <span class="mr-3">
                  <v-btn class="mx-2" fab dark @click="remove(todo.id)" color="red" x-small>
                    <v-icon dark>mdi-delete</v-icon>
                  </v-btn>
                </span>
              </div>
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
        <TodoForm is-new />
      </v-col>
    </v-row>
  </v-content>
</template>

<script>
  import moment from "moment"
  import { formatDate, parseDate } from "../util";
  import TodoForm from "../components/TodoForm";

  export default {
    name: "Todo",
    components: {TodoForm},
    created() {
      this.$store.dispatch('todo/getList');
    },
    computed: {
      getTodoList() {
        return this.$store.getters['todo/getList']
      },
      getError() {
        return this.$store.getters['error/getError']
      }
    },
    methods: {
      formatDate(date) {
        return formatDate(date)
      },
      parseDate(date) {
        return parseDate(date)
      },
      async finish(id) {
        const params = {
          finished_at: moment().format("YYYY-MM-DD HH:mm:ss")
        };

        await this.$store.dispatch('todo/modify', {id, params});
        if (this.getError === '') {
          this.$store.dispatch('todo/getList')
        }
      },
      async remove(id) {
        await this.$store.dispatch('todo/remove', {id});
        if (this.getError === '') {
          this.$store.dispatch('todo/getList')
        }
      }
    },
  }
</script>

<style scoped>
  .todo--header {
    flex-direction: column !important;
  }

  .todo--memo {
    white-space: pre-line;
  }
</style>