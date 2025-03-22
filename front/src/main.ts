import './style.css'

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <form>
        <label for="image_field">sube la imagen amigaaaa</label>
        <input type="file" name="image_field" value="">
        <button type="submit">PICALE</button>
    </form>
  <div id="pics">
    <h1>Fotos</h1>
  </div>
`

const form = document.querySelector("form")

form?.addEventListener('submit', async (e: Event) => {
  e.preventDefault();
  //const server = "http://25.0.119.160:4076/v1/image"
  const server = "http://localhost:4076/v1/image"
  const image_input = document.getElementsByName('image_field')[0] as HTMLInputElement;
  const files = image_input.files
  if (files == null || !files) return alert("bro upload af ile")
  const form_data = new FormData();
  const image = files[0];
  form_data.append("image_field", image)
  fetch(server, {
    method: "POST",
    body: form_data
  })
    .then(async (response) => {
      const json = await response.json()
      const pics = document.getElementById('pics')
      pics?.append(`${json.file_name}`)
    })
    .catch(err => alert(err));
})
