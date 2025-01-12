# Changing the Testkube Context to Cloud

To set the Testkube context to Cloud, the Testkube CLI tool needs to have access, so first you'll need to create an API token with 
valid access rights.

You can create a token:

* with "admin" access rights (access to all environments):

![admin-token](https://user-images.githubusercontent.com/30776/229772185-01f1e466-b04d-4c6d-9d5c-e4464d651177.png)

* with particular role for given environments:

![roles-for-token](https://user-images.githubusercontent.com/30776/229772310-64bda85d-57a8-47b7-a68b-2625089724f8.png)



When your token is there, you're ready to change the Testkube CLI context: 

![setting-context](https://user-images.githubusercontent.com/30776/229771159-4415aa74-70bb-4684-9511-449d0779b483.png)


# Changing the Testkube Context to kubeconfig-based

When you want to return to using Testkube CLI with your local OSS Testkube cluster,  set the context to kubeconfig-based: 

```sh 
testkube set context --kubeconfig
```

