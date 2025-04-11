package fake

import (
	"context"
	"fmt"

	"github.com/argoproj-labs/gitops-promoter/api/v1alpha1"
	ginkov2 "github.com/onsi/ginkgo/v2"
	v1 "k8s.io/api/core/v1"
)

type GitAuthenticationProvider struct {
	scmProvider *v1alpha1.ScmProvider
	secret      *v1.Secret
}

func NewFakeGitAuthenticationProvider(scmProvider *v1alpha1.ScmProvider, secret *v1.Secret) GitAuthenticationProvider {
	return GitAuthenticationProvider{
		scmProvider: scmProvider,
		secret:      secret,
	}
}

func (gh GitAuthenticationProvider) GetGitHttpsRepoUrl(gitRepo v1alpha1.GitRepository) string {
	gitServerPort := 5000 + ginkov2.GinkgoParallelProcess()
	gitServerPortStr := fmt.Sprintf("%d", gitServerPort)

	if gh.scmProvider.Spec.Fake != nil && gh.scmProvider.Spec.Fake.Domain == "" {
		return fmt.Sprintf("http://localhost:%s/%s/%s", gitServerPortStr, gitRepo.Spec.Fake.Owner, gitRepo.Spec.Fake.Name)
	}
	return fmt.Sprintf("http://localhost:%s/%s/%s", gitServerPortStr, gitRepo.Spec.Fake.Owner, gitRepo.Spec.Fake.Name)
}

// Add support for SSH in the Fake provider.
func (gh GitAuthenticationProvider) GetGitRepoUrl(gitRepository v1alpha1.GitRepository) string {
	if gh.scmProvider.Spec.Protocol == "SSH" {
		if gh.scmProvider.Spec.Fake != nil && gh.scmProvider.Spec.Fake.Domain != "" {
			return fmt.Sprintf("git@%s:%s/%s.git", gitServerPortStr, gitRepo.Spec.Fake.Owner, gitRepo.Spec.Fake.Name)
		}
		return fmt.Sprintf("git@localhost:%s/%s.git", gitServerPortStr, gitRepo.Spec.Fake.Owner, gitRepo.Spec.Fake.Name)
	}
	return gh.GetGitHttpsRepoUrl(gitRepository)
}

func (gh GitAuthenticationProvider) GetToken(ctx context.Context) (string, error) {
	return "", nil
}

func (gh GitAuthenticationProvider) GetUser(ctx context.Context) (string, error) {
	return "git", nil
}
