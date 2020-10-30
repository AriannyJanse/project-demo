<template>
  <v-data-table :headers="headers" :items="users" class="elevation-1">
    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title>USERS</v-toolbar-title>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="500px">
          <template v-slot:activator="{ on, attrs }">
            <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on">
              New User
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-form v-model="valid">
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedUser.nickname"
                        :rules="[rules.required]"
                        label="Nickname"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedUser.email"
                        :rules="[rules.required ,rules.email]"
                        label="Email"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedUser.password"
                        label="Password"
                        :rules="[rules.required]"
                        type="password"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model.number="editedUser.company_id"
                        :rules="[rules.required ,rules.companyID]"
                        label="Company ID"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-form>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="close">
                Cancel
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                @click="save"
                :disabled="!valid"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="dialogDelete" max-width="500px">
          <v-card>
            <v-card-title class="headline">
              Are you sure you want to delete this user?
            </v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="closeDelete"
                >Cancel</v-btn
              >
              <v-btn color="blue darken-1" text @click="deleteUserConfirm"
                >OK</v-btn
              >
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="alertDialog" width="300">
          <v-alert
            style="margin-bottom: 0;"
            type="error"
            transition="scale-transition"
          >
            {{ alertMessage }}
          </v-alert>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:[`item.password`]="{}">
      *********
    </template>
    <template v-slot:[`item.actions`]="{ item }">
      <v-icon small class="mr-2" @click="editUser(item)">
        fas fa-edit
      </v-icon>
      <v-icon small @click="deleteUser(item)">
        fas fa-trash-alt
      </v-icon>
    </template>
    <template v-slot:no-data>
      <v-btn color="primary" @click="initialize">
        Reset
      </v-btn>
    </template>
  </v-data-table>
</template>

<script>
import api from "@/api";
export default {
  data: () => ({
    valid: false,
    alertDialog: false,
    alertMessage: "",
    dialog: false,
    dialogDelete: false,
    headers: [
      {
        text: "Nickname",
        align: "start",
        sortable: false,
        value: "nickname"
      },
      { text: "Email", value: "email" },
      { text: "Password", value: "password" },
      { text: "Company ID", value: "company_id" },
      { text: "Actions", value: "actions", sortable: false }
    ],
    users: [],
    editedIndex: -1,
    editedUser: {
      nickname: "",
      email: "",
      password: "",
      company_id: 0
    },
    defaultUser: {
      nickname: "",
      email: "",
      password: "",
      company_id: 0
    },
    rules: {
      required: value => !!value || "This field is required",
      email: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        return pattern.test(value) || "Invalid email.";
      },
      companyID: value => value > 0 || "Company ID must be greater than 0"
    }
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "New User" : "Edit User";
    }
  },

  watch: {
    dialog(value) {
      value || this.close();
    },
    dialogDelete(value) {
      value || this.closeDelete();
    }
  },

  created() {
    this.initialize();
  },

  methods: {
    initialize() {
      return api.getUsers().then(response => {
        response.data.status == 500
          ? (this.alertDialog = true) &&
            (this.alertMessage = response.data.message)
          : (this.users = response.data);
      });
    },

    editUser(item) {
      this.editedIndex = this.users.indexOf(item);
      this.editedUser = Object.assign({}, item);
      this.dialog = true;
    },

    deleteUser(item) {
      this.editedIndex = this.users.indexOf(item);
      this.editedUser = Object.assign({}, item);
      this.dialogDelete = true;
    },

    deleteUserConfirm() {
      return api.deleteUser(this.editedUser.ID).then((response) => {
        response.status == 400 || response.data.status == 500
          ? (this.alertDialog = true) &&
            (this.alertMessage = response.data.message)
          : this.users.splice(this.editedIndex, 1) && this.closeDelete();
      });
    },

    close() {
      this.dialog = false;
      this.$nextTick(() => {
        this.editedUser = Object.assign({}, this.defaultUser);
        this.editedIndex = -1;
      });
    },

    closeDelete() {
      this.dialogDelete = false;
      this.$nextTick(() => {
        this.editedUser = Object.assign({}, this.defaultUser);
        this.editedIndex = -1;
      });
    },

    save() {
      if (this.editedIndex > -1) {
        return api
          .updateUser(this.editedUser.ID, this.editedUser)
          .then((response) => {
            response.status == 400 || response.data.status == 500
              ? (this.alertDialog = true) &&
                (this.alertMessage = response.data.message)
              : Object.assign(this.users[this.editedIndex], this.editedUser) &&
                this.close();
          });
      } else {
        return api.createUser(this.editedUser).then((response) => {
          response.data.status == 400 || response.data.status == 500
            ? (this.alertDialog = true) &&
              (this.alertMessage = response.data.message)
            : this.users.push(this.editedUser) && this.close();
        });
      }
    }
  }
};
</script>