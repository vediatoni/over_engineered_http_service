## Service url: https://app-oehs-qs4uvgfbba-uc.a.run.app
## Tech stack 
**Go v1.17.3**
the main http service is coded in Go.

**Google Cloud Platform - GCP**
for infrastructure: 
 - [Cloud Build for CI/CD](https://www.google.com/search?q=cloud%20build&oq=cloud%20build&aqs=edge..69i57j69i59l4j0i512j69i60l3.2358j0j4&sourceid=chrome&ie=UTF-8) 
 - [Cloud Run for running the service](https://cloud.google.com/run) 
 - [Artifact Registry for storing container images](https://cloud.google.com/artifact-registry)
 - [Kaniko for building container images via Cloud Build](https://github.com/GoogleContainerTools/kaniko)

**Terraform**
for infrastructure provisioning on GCP

**Dockerfile**
for packaging app into container

## Project file structure
 **/cmd**
main code for this project.

**/build**
packaging and Cloud Build file for CI/CD

**/deployments**
terraform files

[You can read more about the structure here](https://github.com/golang-standards/project-layout)

## Instructions
1. Create a project on GCP  
2. Create a service account on GCP (IAM) and download json auth keys for this account *(make sure to give it the right permissions- **Editor, Project IAM Admin**)*  
3. Move the json auth file to the root path of this project 
4. Connect repository to the Cloud Build > trigger  
5. Create `deployments/terraform.tfvars` file and make sure to configure these variables: *artifact_registry, project, organization, credentials_file* use `deployments/variables.tf` as a reference
7. Run terraform apply  
8. Run Cloud Build trigger (via GUI or git push event)