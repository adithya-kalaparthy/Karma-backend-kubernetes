# Karma-backend-kubernetes

Backend code for karma application deployed using kubernetes on EKS cluster.

### Pre-requistes:

- kubectl (Installed along with minikube)
- [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [eksctl](https://eksctl.io/installation/)

### Production setup:
- Follow the steps laid out in [this file](https://www.notion.so/AWS-EKS-setup-16ffd5111c0d80729fb8cf85083d8957)

### Local setup:
- Follow the steps laid out in [local deployment file](https://www.notion.so/K8s-Local-setup-183fd5111c0d80979e5bd7c183c45d72)

### Repo project structure:
- Click [here](https://www.notion.so/K8s-EKS-project-setup-82aaab6363fc465cbe42f24aa5fe08ee) to view the structure.

### Deployment:
- Deployment done with the help of [this yt video](https://www.youtube.com/watch?v=9qSmFWwsxwA)
- All the production deployment files are stored in k8s-prod folder.
- If you push to main branch, github actions will create a new artifact in karma-api-k8s repository.
- It will be tagged with commit sha and production-latest tags.
- Since the node instance type in our cluster is t3.micro, it cannot handle rollout deployments.
- Hence the pod replicas for karma-deployment will be made to 0.
- And then, all the necessary services, ingresses, and deployments will be done.

