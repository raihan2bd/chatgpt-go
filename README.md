# NeuralChat - using Go (golang) and OpenAI (ChatGPT)
<p>NeuralChat is a cutting-edge chatbot project that allows users to chat with the state-of-the-art OpenAI chatbot, GPT-3.5, in a seamless and user-friendly manner. The project is built using Golang, PostgreSQL, Javascript, HTML, and CSS, making it a powerful and robust solution for all your chatbot needs.

NeuralChat implements stateful authentication, which ensures that users can securely log in and engage with the chatbot. The user experience is smooth and intuitive, thanks to the project's sleek and modern design.

Using NeuralChat, users can interact with GPT-3.5 through prompts, allowing them to chat about a wide range of topics and receive intelligent and insightful responses. The chatbot is trained to understand natural language, making conversations feel human-like and engaging.

Overall, NeuralChat is a fantastic solution for anyone looking to build a powerful chatbot that can engage with users in a natural and intuitive way. With its stateful authentication and advanced GPT-3.5 technology, this project is sure to impress and delight users.</p>

### Tech Stack <a name="tech-stack"></a>

I used Go (golang), PostgreSQL, OpenAI, Html, Javascript and CSS to build this project.
  <summary>Full Stack</summary>
  <ul>
    <li>Go (Golang)</li>
    <li>PostgreSQL</li>
    <li>OpenAI</li>
    <li>JAVASCRIPT</li>
    <li>Html</li>
    <li>CSS</li>
  </ul>

 <summary>Major Dependencies</summary>
  <ul>
    <li><a href="https://github.com/go-chi/chi">Chi router</a></li>
    <li><a href="https://github.com/alexedwards/scs/v2">Alex edwards SCS </a> Session Manager</li>
    <li><a href="https://github.com/jackc/pgx">PGX PostgreSQL Driver</a> Database Driver</li>
    <li><a href="https://github.com/justinas/nosurf">Nosurf</a> Generate CSRFToken</li>
    <li><a href="https://github.com/sashabaranov/go-openai">OpenAI</a> OpenAI Driver</li>
  </ul>

## Demo
![Capture_home](https://user-images.githubusercontent.com/35267447/229288848-3b920169-b580-4121-b0c6-cc431bb12fc6.PNG)

![Capture](https://user-images.githubusercontent.com/35267447/229288865-4947e6e6-bfce-4328-8c03-f0a37f542039.PNG)

![Capture_home_m](https://user-images.githubusercontent.com/35267447/229288832-b06b36aa-6b4d-432f-87a3-86bf4676005f.PNG)
![Capture-m](https://user-images.githubusercontent.com/35267447/229288895-d0b898ca-5470-47af-9e48-b7644e10c665.PNG)



## 💻 Getting Started
- To get star with this package first of all you have to clone the project ⬇️
``` bash
https://github.com/raihan2bd/chatgpt-go.git
```
- Then Make sure you have install [Go (golang)](https://go.dev/dl/) version 1.8.0 or latest stable version.
- Then make sure you have install [PostgreSQL](https://www.postgresql.org/) on your local machine if you want to use this project as locally.
- Then Create a database called `chatgpt` inside the database create two table with below command ⬇️

```sql
 CREATE TABLE public.users (
	id serial4 NOT NULL,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(60) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	access_level int4 NULL DEFAULT 1,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE sessions (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```
- To install all the Go packages navigate the folder address on your terminal and enter the below command ⬇️
``` bash
go mod tidy
```

![Capture](https://user-images.githubusercontent.com/35267447/229288804-fc4769a8-6e29-4c0b-b2fa-750d27fda43b.PNG)
- Make sure you rename `example.env` file to `.env` file and modify it with your database credentials (you created earlier) and other info too.


# Usages
> *Note: Before enter the below command make sure you are in the right directory.*

- After finishing the above instructions you can see the project in your local machine by entering the below command ⬇️
```bash
go run main.go
```

- Then you can see this project live on your browser by this link http://localhost:8080 or your given the port number you set for the project.


## 👥 Author

👤 **Abu Raihan**

- GitHub: [@githubhandle](https://github.com/raihan2bd)
- Twitter: [@twitterhandle](https://twitter.com/raihan2bd)
- LinkedIn: [LinkedIn](https://linkedin.com/in/raihan2bd)

## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thanks [Trevor Sawler](https://www.gocode.ca/) Who help me a lot learn go.
I would like to thanks [OpenAI](https://openai.com/) for providing us this awesome and excitant API 🙏.

## ⭐️ Show your support <a name="support"></a>

Thanks for visiting my repository. Give a ⭐️ if you like this project!

## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

## Contribution
*Your suggestions will be more than appreciated. If you want to suggest anything for this project feel free to do that. :slightly_smiling_face:*
