<template>
  <div class="container mt-5">
    <h1 class="text-center mb-4">Cadastro de Alunos</h1>
    <AlunoForm @alunoCadastrado="fetchAlunos" />
    <h2 class="mb-3">Lista de Alunos</h2>
    <AlunoList :alunos="alunos" />
  </div>
</template>

<script>
import AlunoForm from './components/AlunoForm.vue';
import AlunoList from './components/AlunoList.vue';

export default {
  components: {
    AlunoForm,
    AlunoList
  },
  data() {
    return {
      alunos: []
    };
  },
  methods: {
    async fetchAlunos() {
      try {
        const response = await fetch('http://localhost:8080/aluno');
        const data = await response.json();
        this.alunos = data;
      } catch (error) {
        console.error('Erro ao buscar alunos:', error);
      }
    }
  },
  mounted() {
    this.fetchAlunos();
  }
};
</script>
