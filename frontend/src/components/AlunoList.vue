<template>
  <div>
    <!-- Lista de Alunos -->
    <ul id="aluno-list" class="list-group">
      <li
          v-for="aluno in alunos"
          :key="aluno.ID"
          class="list-group-item d-flex justify-content-between align-items-center"
          @click="openModal(aluno.ID)"
      >
        <span><strong>NOME:</strong> {{ aluno.nome }}, <strong>TELEFONE:</strong> {{ aluno.telefone }}</span>
      </li>
    </ul>

    <!-- Modal para editar aluno -->
    <div
        class="modal"
        id="alunoEditModal"
        tabindex="-1"
        role="dialog"
        aria-labelledby="alunoEditModalLabel"
        aria-hidden="true"
    >
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="alunoEditModalLabel">Editar Aluno</h5>
            <button
                type="button"
                class="close"
                data-dismiss="modal"
                aria-label="Close"
            >
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="salvarAluno">
              <div class="form-group">
                <label for="cpf">CPF:</label>
                <input
                    type="text"
                    v-model="aluno.cpf"
                    id="cpf"
                    name="cpf"
                    class="form-control"
                    required
                />
              </div>
              <div class="form-group">
                <label for="estado">Estado:</label>
                <input
                    type="text"
                    v-model="aluno.estado"
                    id="estado"
                    name="estado"
                    class="form-control"
                    required
                />
              </div>
              <div class="form-group">
                <label for="cidade">Cidade:</label>
                <input
                    type="text"
                    v-model="aluno.cidade"
                    id="cidade"
                    name="cidade"
                    class="form-control"
                    required
                />
              </div>
              <div class="form-group">
                <label for="bairro">Bairro:</label>
                <input
                    type="text"
                    v-model="aluno.bairro"
                    id="bairro"
                    name="bairro"
                    class="form-control"
                    required
                />
              </div>
              <div class="form-group">
                <label for="rua">Rua:</label>
                <input
                    type="text"
                    v-model="aluno.rua"
                    id="rua"
                    name="rua"
                    class="form-control"
                    required
                />
              </div>
              <div class="form-group">
                <label for="descricao">Descrição:</label>
                <textarea
                    v-model="aluno.descricao"
                    id="descricao"
                    name="descricao"
                    class="form-control"
                ></textarea>
              </div>
              <button type="submit" class="btn btn-primary">
                Salvar Alterações
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/* global $ */

export default {
  props: {
    alunos: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      aluno: {
        id: null,
        nome: '',
        cpf: '',
        estado: '',
        cidade: '',
        bairro: '',
        rua: '',
        descricao: '',
      },
    };
  },
  methods: {
    openModal(id) {
      console.log('Tentando abrir o modal com ID:', id);
      this.aluno.id = id;
      // Você pode adicionar lógica aqui para buscar os dados do aluno específico antes de abrir o modal, se necessário
      $('#alunoEditModal').modal('show');
    },
    async salvarAluno() {
      try {
        const response = await fetch(`http://localhost:8080/alunoEdit/${this.aluno.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.aluno),
        });

        if (response.ok) {
          alert('Aluno atualizado com sucesso!');
          this.$emit('alunoCadastrado');
          $('#alunoEditModal').modal('hide');
        } else {
          alert('Erro ao atualizar aluno.');
        }
      } catch (error) {
        console.error('Erro ao atualizar aluno:', error);
      }
    },
  },
};
</script>

<style scoped>
/* Estilo básico para a lista de alunos */
ul#aluno-list {
  margin-top: 20px;
  padding-left: 0;
  list-style: none;
}

.list-group-item {
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.list-group-item:hover {
  background-color: #f8f9fa;
}

/* Estilo do modal */
.modal-header {
  background-color: #007bff;
  color: white;
}

.modal-title {
  font-weight: bold;
}

.modal-body {
  padding: 20px;
}

.modal-footer {
  padding: 10px;
  text-align: right;
}

.btn-primary {
  background-color: #007bff;
  border-color: #007bff;
  transition: background-color 0.3s ease;
}

.btn-primary:hover {
  background-color: #0056b3;
  border-color: #004085;
}

/* Estilo para os campos do formulário no modal */
.form-group label {
  font-weight: bold;
}

.form-group input,
.form-group textarea {
  border-radius: 4px;
  border: 1px solid #ced4da;
  padding: 10px;
}

.form-group input:focus,
.form-group textarea:focus {
  border-color: #80bdff;
  outline: none;
  box-shadow: 0 0 5px rgba(0, 123, 255, 0.5);
}
</style>