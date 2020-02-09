<template>
  <v-content>
    <v-container>
      <v-row class="text-center" align="center" justify="center">
        <v-col class="mb-4" cols="12" sm="8" md="4">
          <v-toolbar color="primary" dark>
            <v-toolbar-title>会員登録</v-toolbar-title>
          </v-toolbar>
          <v-card>
            <v-form>
              <v-col>
                <v-text-field
                    type="text"
                    v-model="username"
                    label="ユーザー名"
                    :error-messages="usernameErrors"
                    required
                    @input="$v.username.$touch()"
                    @blur="$v.username.$touch()"
                />
              </v-col>
              <v-col>
                <v-text-field
                    type="email"
                    v-model="email"
                    label="メールアドレス"
                    :error-messages="emailErrors"
                    required
                    @input="$v.email.$touch()"
                    @blur="$v.email.$touch()"
                />
              </v-col>
              <v-col>
                <v-text-field
                    type="password"
                    v-model="password"
                    label="パスワード"
                    :error-messages="passwordErrors"
                    required
                    @input="$v.password.$touch()"
                    @blur="$v.password.$touch()"
                />
              </v-col>
              <v-col>
                <v-btn color="primary" @click="login">会員登録</v-btn>
              </v-col>
            </v-form>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
  import { validationMixin } from 'vuelidate'
  import { email, minLength, required } from 'vuelidate/lib/validators'

  export default {
    name: "Login",
    mixins: [validationMixin],
    validations: {
      username: {required},
      email: {required, email},
      password: {required, minLength: minLength(8)},
    },
    data() {
      return {
        username: '',
        email: '',
        password: '',
      }
    },
    computed: {
      usernameErrors() {
        const errors = [];
        if (!this.$v.username.$dirty) return errors;
        !this.$v.username.required && errors.push('ユーザー名を入力してください');
        return errors;
      },
      emailErrors() {
        const errors = [];
        if (!this.$v.email.$dirty) return errors;
        !this.$v.email.required && errors.push('メールアドレスを入力してください');
        !this.$v.email.email && errors.push('メールアドレスの形式が不正です');
        return errors;
      },
      passwordErrors() {
        const errors = [];
        if (!this.$v.password.$dirty) return errors;
        !this.$v.password.required && errors.push('パスワードを入力してください');
        !this.$v.password.minLength && errors.push('パスワードは8文字以上で入力してください');
        return errors;
      }
    },
    methods: {
      login() {
        this.$v.$touch();
        if (this.$v.$invalid) return;

        alert("sending")
      }
    }
  }
</script>

<style scoped>

</style>