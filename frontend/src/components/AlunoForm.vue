<template>
  <form @submit.prevent="cadastrarAluno" class="mb-4 p-3 bg-light rounded shadow-sm">
    <div class="form-group">
      <label for="nome">Nome:</label>
      <input type="text" v-model="nome" id="nome" name="nome" class="form-control" required>
    </div>
    <div class="form-group">
      <label for="telefone">Telefone:</label>
      <input type="text" v-model="telefone" id="telefone" name="telefone" class="form-control" required>
    </div>
    <button type="submit" class="btn btn-primary btn-block mt-3">Cadastrar</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      nome: '',
      telefone: ''
    };
  },
  methods: {
    async cadastrarAluno() {
      try {
        await fetch('http://localhost:8080/aluno', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({nome: this.nome, telefone: this.telefone})
        });

        alert('Aluno cadastrado com sucesso!');
        this.$emit('alunoCadastrado');
        this.nome = '';
        this.telefone = '';
      } catch (error) {
        console.error('Erro ao cadastrar aluno:', error);
        alert('Erro ao cadastrar aluno.');
      }
    }
  }
};
</script>
