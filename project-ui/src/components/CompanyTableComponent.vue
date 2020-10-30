<template>
  <v-data-table :headers="headers" :items="companies" class="elevation-1">
    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title>Companies</v-toolbar-title>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="500px">
          <template v-slot:activator="{ on, attrs }">
            <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on">
              New Company
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                  <v-col cols="12" sm="6" md="4">
                    <v-text-field
                      v-model="editedCompany.name"
                      label="Name"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="close">
                Cancel
              </v-btn>
              <v-btn color="blue darken-1" text @click="save">
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="dialogDelete" max-width="550px">
          <v-card>
            <v-card-title class="headline">
              Are you sure you want to delete this company?
            </v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="closeDelete"
                >Cancel</v-btn
              >
              <v-btn color="blue darken-1" text @click="deleteCompanyConfirm"
                >OK</v-btn
              >
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:[`item.actions`]="{ item }">
      <v-icon small class="mr-2" @click="editCompany(item)">
        fas fa-edit
      </v-icon>
      <v-icon small @click="deleteCompany(item)">
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
    dialog: false,
    dialogDelete: false,
    headers: [
      {
        text: "Name",
        align: "start",
        value: "name"
      },
      { text: "Actions", value: "actions", sortable: false }
    ],
    companies: [],
    editedIndex: -1,
    editedCompany: {
      name: ""
    },
    defaultCompany: {
      name: ""
    }
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "New Company" : "Edit Company";
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
      return api.getCompanies().then(response => {
        this.companies = response.data;
      });
    },

    editCompany(item) {
      this.editedIndex = this.companies.indexOf(item);
      this.editedCompany = Object.assign({}, item);
      this.dialog = true;
    },

    deleteCompany(item) {
      this.editedIndex = this.companies.indexOf(item);
      this.editedCompany = Object.assign({}, item);
      this.dialogDelete = true;
    },

    deleteCompanyConfirm() {
      return api.deleteCompany(this.editedCompany.ID).then(() => {
        this.companies.splice(this.editedIndex, 1) && this.closeDelete();
      });
    },

    close() {
      this.dialog = false;
      this.$nextTick(() => {
        this.editedCompany = Object.assign({}, this.defaultCompany);
        this.editedIndex = -1;
      });
    },

    closeDelete() {
      this.dialogDelete = false;
      this.$nextTick(() => {
        this.editedCompany = Object.assign({}, this.defaultCompany);
        this.editedIndex = -1;
      });
    },

    save() {
      if (this.editedIndex > -1) {
        return api.updateCompany(this.editedCompany.ID, this.editedCompany).then(() => {
          Object.assign(this.companies[this.editedIndex], this.editedCompany) &&
            this.close();
        });
      } else {
        return api.createCompany(this.editedCompany).then(() => {
          this.companies.push(this.editedCompany) && this.close();
        });
      }
    }
  }
};
</script>