<template>
  <v-content>
    <v-row class="text-center" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <h1>TodoList</h1>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-switch v-model="isShowFinished" :label="`完了済みを表示${isShowFinished ? 'しない' : 'する'}`"/>
    </v-row>
    <div v-if="this.isLoadingOn">
      <v-row justify="center">
        <v-progress-circular :size="50" color="primary" indeterminate/>
      </v-row>
    </div>
    <div v-else>
      <v-row v-if="this.error !== ''" class="text-center" justify="center">
        <v-alert type="error">{{ this.error }}</v-alert>
      </v-row>
      <v-row justify="center">
        <v-col class="mb-3" cols="12" xs="12" sm="8">
          <v-expansion-panels class="todo-container" multiple>
            <v-expansion-panel v-for="todo in this.todoList" :key="todo.id">
              <v-expansion-panel-header :disable-icon-rotate="!!todo.finished_at">
                <div class="d-flex todo--header">
                  <div class="mb-3">
                    <span class="mr-3"><i class="fas fa-stopwatch mr-3"/>{{ formatDate(todo.limit_date) }}</span>
                  </div>
                  <h2 class="todo--title">{{ todo.title }}</h2>
                </div>
                <template v-if="todo.finished_at" v-slot:actions>
                  <v-icon color="teal">mdi-check</v-icon>
                </template>
              </v-expansion-panel-header>
              <v-expansion-panel-content class="todo--memo">
                {{ todo.memo }}
                <div class="text-right">
                <span class="mr-3">
                  <v-btn v-if="!todo.finished_at" class="mx-2" fab dark @click="finish(todo.id)" color="primary"
                         x-small>
                    <v-icon dark>mdi-check</v-icon>
                  </v-btn>
                   <v-btn v-else class="mx-2" fab dark @click="unFinished(todo.id)" color="secondary"
                          x-small>
                    <v-icon dark>mdi-backspace</v-icon>
                  </v-btn>
                </span>
                  <span class="mr-3">
                  <TodoForm
                      :todo-id="todo.id"
                      :todo-title="todo.title"
                      :todo-memo="todo.memo"
                      :todo-limit-date="todo.limit_date"
                      @get-todo-list="getTodoList"
                      @handle-error="handleError"
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
          <TodoForm @get-todo-list="getTodoList" @handle-error="handleError" is-new/>
        </v-col>
      </v-row>
    </div>
  </v-content>
</template>

<script>
import { formatDate, generateTitle, parseDate, STATUS_UNAUTHORIZED } from "@/util";
import TodoForm from "../components/TodoForm";

export default {
  name: "Todo",
  title: generateTitle('TodoList'),
  components: {TodoForm},
  data() {
    return {
      isShowFinished: false
    }
  },
  created() {
    this.getTodoList();
  },
  watch: {
    isShowFinished() {
      this.getTodoList()
    }
  },
  computed: {
    todoList() {
      return this.$store.getters['todo/getList']
    },
    isLoadingOn() {
      return this.$store.getters['loading/isOn']
    },
    error() {
      return this.$store.getters['error/getError']
    },
    errorCode() {
      return this.$store.getters['error/getCode']
    }
  },
  methods: {
    getTodoList() {
      const params = {
        is_show_finished: this.isShowFinished
      };

      (async () => {
        await this.$store.dispatch('todo/getList', params);
        if (this.error !== '') {
          return this.handleError();
        }
      })()
    },
    formatDate(date) {
      return formatDate(date)
    },
    parseDate(date) {
      return parseDate(date)
    },
    async finish(id) {
      await this.$store.dispatch('todo/finished', {id});
      if (this.error !== '') {
        return this.handleError();
      }
      this.getTodoList();
    },
    async unFinished(id) {
      await this.$store.dispatch('todo/unfinished', {id});
      if (this.error !== '') {
        return this.handleError();
      }
      this.getTodoList();
    },
    async remove(id) {
      await this.$store.dispatch('todo/remove', {id});
      if (this.error !== '') {
        return this.handleError();
      }
      this.getTodoList();
    },
    handleError() {
      if (this.errorCode === STATUS_UNAUTHORIZED) {
        this.$store.dispatch('auth/logout');
      }
    }
  },
}
</script>

<style scoped>
.todo-container {
  margin-left: 5px;
}

.todo--header {
  flex-direction: column !important;
}

.todo--title {
  overflow-wrap: anywhere;
}

.todo--memo {
  white-space: pre-line;
  overflow-wrap: anywhere;
}
</style>