<script>
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  // import Counter from './lib/Counter.svelte'
  let texto = ""
  let valorSalida = ""

  const onPostAsync = async () => {
    var puerto = 9700
    try {
      const response = await fetch(`http://localhost:${puerto}/commands`, 
        {
          method: "POST",
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Comandos: texto
          })
        });
  
      if (!response.ok) throw new Error(`Error en la peticion`)
      
      var salida = await response.json()
      valorSalida = salida
    } 
    catch (ex) {
      console.log(`Error: ${ex}`)
    }
  }

  function cargaArchivo(e) {
    // document.getElementById("inputArchivo").click()
    const file = e.target.files[0]
    if (!file) return

    // const nombre = new file.name
    const lector = new FileReader()

    lector.onload = (temp) => {
      texto = temp.target.result.toString()
    }

    lector.readAsText(file, "UTF-8")
  }
</script>

<main>
  <form  method="post">
    
    <div id="boton">
      <input
        type="file"
        id="inputArchivo"
        accept=".asdj"
        style="display:none"
        on:change={cargaArchivo}
      />
      <label for="inputArchivo">
        Carga
      </label>

      <button on:click={onPostAsync}>
        Enviar
      </button>
    </div>
    <div id="recuadro">
      <textarea id="area-texto" bind:value={texto}>
        
      </textarea>
      <textarea name="salida" id="salida" disabled>

      </textarea>
    </div>
  </form>

</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
