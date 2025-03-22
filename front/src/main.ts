import './style.css'

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
<div>
  <img src="https://media2.giphy.com/media/YTtqB2j5EN7IA/200w.gif?cid=6c09b952avx9x0b6g87aitmsno4dk223nzhkjc4wj6naz1ye&ep=v1_gifs_search&rid=200w.gif&ct=g"/>
        <label for="ip">ip del server d minecraft hipixel</label>
        <input type="text" id="ip" value="">
  </div>
  <form>
        <label for="image_field">sube la imagen amigaaaa</label>
        <input type="file" name="image_field" value="">
        <button type="submit">PICALE</button>
    </form>
  <div id="pics">
    <span>Fotos</span>
  </div>
`

const form = document.querySelector("form")

form?.addEventListener('submit', async (e: Event) => {
  e.preventDefault();
  const ip = document.querySelector("#ip") as HTMLInputElement
  const server = `http://${ip?.value}:4076/v1/image`
  //const server = "http://localhost:4076/v1/image"
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
      const span: HTMLSpanElement = document.createElement('span');
      span.innerText = json.file_name
      pics?.appendChild(span)
    })
    .catch(err => alert(err));
})
