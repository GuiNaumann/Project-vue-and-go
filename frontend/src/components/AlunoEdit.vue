<template>
  <div>
    <h2>Editar Aluno</h2>
    <form @submit.prevent="salvarAluno">
      <div class="form-group">
        <label for="nome">Nome:</label>
        <input type="text" v-model="aluno.nome" id="nome" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="cpf">CPF:</label>
        <input type="text" v-model="aluno.cpf" id="cpf" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="estado">Estado:</label>
        <input type="text" v-model="aluno.estado" id="estado" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="cidade">Cidade:</label>
        <input type="text" v-model="aluno.cidade" id="cidade" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="bairro">Bairro:</label>
        <input type="text" v-model="aluno.bairro" id="bairro" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="rua">Rua:</label>
        <input type="text" v-model="aluno.rua" id="rua" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="descricao">Descrição:</label>
        <textarea v-model="aluno.descricao" id="descricao" class="form-control"></textarea>
      </div>
      <button type="submit" class="btn btn-primary">Salvar Alterações</button>
    </form>
  </div>
</template>

<script>
export default {
  props: {
    id: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      aluno: {
        nome: '',
        cpf: '',
        estado: '',
        cidade: '',
        bairro: '',
        rua: '',
        descricao: ''
      }
    };
  },
  mounted() {
    this.buscarAluno();
  },
  methods: {
    async buscarAluno() {
      try {
        const response = await fetch(`http://localhost:8080/alunoEdit/${this.id}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        });

        const data = await response.json();
        this.aluno = data;
      } catch (error) {
        console.error('Erro ao buscar os dados do aluno:', error);
      }
    },
    async salvarAluno() {
      try {
        const response = await fetch(`http://localhost:8080/alunoEdit/${this.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.aluno)
        });

        if (response.ok) {
          alert('Aluno atualizado com sucesso!');
          this.$emit('alunoAtualizado');
        } else {
          alert('Erro ao atualizar aluno.');
        }
      } catch (error) {
        console.error('Erro ao atualizar o aluno:', error);
      }
    }
  }
};
</script>

<style scoped>
.form-group {
  margin-bottom: 15px;
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
</style>
