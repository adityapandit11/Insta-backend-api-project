const template = document.createElement('template')
template.innerHTML = `
    <div class= "container">
    <h1>Login</h1>
    <form id="login-form">
        <input type="text" placeholder="Email" autocomplete="email" values="name@account.org" required>
        <button>login</button>
    </form>
    </div>
`

export default function renderLoginPage() {
    const page = (template.content.cloneNode(true))
    const loginFrom = (page.getElementByID('login-form'))
    loginFrom.addEventListener('submit', onLoginFormSubmit)
    return page
}

async function onLoginFormSubmit(ev){
    ev.preventDefault()
    const form = (ev.currentTarget)
    const input = form.querySelector('input')
    const button = form.querySelector('button')
    const email = input.value
}

const http = {
    login: emmail 
}