package main

import (
	"context"
	"fmt"
	"log"
	"time"

	article "github.com/maryam-nokohan/go-article/proto"
	"google.golang.org/grpc"
)

func DoClientStreaming(c article.ArticleServiceClient) {
	fmt.Println("Client streaming ...(writing articles to DB)...")
	stream, err := c.ProcessArticle(context.Background())
	if err != nil {
		log.Fatal(err)
	}

requests := []*article.ArticleRequest{
	{
		Article: &article.Article{
			Title: "The Evolution of Quantum Computing",
			Body: `Quantum computing leverages qubits to perform computations beyond classical capabilities.
			Quantum algorithms like Shor's and Grover's can solve complex problems faster.
			Researchers are exploring quantum supremacy for cryptography and optimization challenges.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Blockchain Beyond Cryptocurrency",
			Body: `Blockchain technology provides secure, decentralized ledgers.
			Smart contracts and decentralized applications enable automation across industries.
			Supply chain, healthcare, and finance are rapidly adopting blockchain solutions.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Advancements in Renewable Energy",
			Body: `Solar, wind, and hydroelectric power are reshaping global energy production.
			Energy storage solutions like lithium-ion batteries improve grid reliability.
			Governments and companies are investing in sustainable technologies to combat climate change.`,
		},
	},
	{
		Article: &article.Article{
			Title: "The Internet of Things (IoT) Revolution",
			Body: `IoT connects devices to the internet, enabling smart homes, cities, and industries.
			Data collected from sensors drives automation and predictive maintenance.
			Edge computing improves performance and reduces latency in IoT systems.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Virtual Reality and Augmented Reality",
			Body: `VR immerses users in digital environments, while AR overlays information onto reality.
			Applications include gaming, education, healthcare, and training simulations.
			Advances in hardware and software are making immersive experiences more accessible.`,
		},
	},
	{
		Article: &article.Article{
			Title: "5G Networks and Connectivity",
			Body: `5G offers high-speed, low-latency communication for mobile and IoT devices.
			Enhanced bandwidth enables applications like autonomous vehicles and remote surgery.
			Telecom providers are rapidly deploying 5G infrastructure worldwide.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Ethical Implications of AI",
			Body: `AI raises questions about privacy, bias, and accountability.
			Fairness in algorithmic decision-making is critical in sectors like hiring and healthcare.
			Regulations and ethical frameworks are emerging to guide responsible AI use.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Advances in Robotics and Automation",
			Body: `Robots are increasingly performing tasks in manufacturing, healthcare, and logistics.
			Collaborative robots work alongside humans safely and efficiently.
			AI-powered automation reduces errors and improves productivity across industries.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Edge Computing and Data Processing",
			Body: `Edge computing processes data closer to the source, reducing latency and bandwidth use.
			It is crucial for real-time applications like autonomous vehicles and smart grids.
			Cloud and edge hybrid architectures provide flexibility and scalability.`,
		},
	},
	{
		Article: &article.Article{
			Title: "The Role of Biotechnology in Medicine",
			Body: `Biotechnology enables personalized medicine and genetic therapies.
			CRISPR and gene editing technologies are advancing treatment of genetic disorders.
			Biotech innovations are transforming drug discovery and disease prevention.`,
		},
	},
	{
		Article: &article.Article{
			Title: "Space Exploration and Technology",
			Body: `Private companies and governments are pushing the boundaries of space travel.
			Satellites, rovers, and telescopes improve our understanding of the universe.
			Innovations in propulsion and materials are making interplanetary missions more feasible.`,
		},
	},
}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server response tags:")
	for _, t := range resp.GetTags() {
		fmt.Printf(" - %s : %d\n", t.Word, t.Freq)
	}
}

func DoUnary(c article.ArticleServiceClient) {
	fmt.Println("Unary call ...(top tags)")

	req := &article.TopTagsRequst{
		Topn: 3,
	}

	res, err := c.TopTags(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top tags:")
	for _, t := range res.GetTags() {
		fmt.Printf(" - %s : %d\n", t.Word, t.Freq)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := article.NewArticleServiceClient(conn)

	DoClientStreaming(client)
	DoUnary(client)
}