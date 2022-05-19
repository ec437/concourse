package topgun_test

import (
	. "github.com/concourse/concourse/topgun/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("A pipeline-provided resource type", func() {
	BeforeEach(func() {
		Deploy("deployments/concourse.yml", "-o", "operations/no-gc.yml")
	})

	It("does not result in redundant containers when running resource actions", func() {
		By("setting a pipeline")
		Fly.Run("set-pipeline", "-n", "-c", "pipelines/custom-types.yml", "-p", "pipe")
		Fly.Run("unpause-pipeline", "-p", "pipe")

		By("triggering the build")
		buildSession := Fly.Start("trigger-job", "-w", "-j", "pipe/get-10m")
		<-buildSession.Exited
		Expect(buildSession.ExitCode()).To(Equal(1))

		By("expecting a container for the resource check, resource type check, task resource type check and task image check")
		Expect(ContainersBy("type", "check")).To(HaveLen(4))

		By("expecting a container for resource type check, resource type get, resource check, resource get in get step. expecting a container for nested resource type check, image check, image get and task run in task step. In total 8 containers.")
		expectedContainersBefore := 8
		Expect(FlyTable("containers")).Should(HaveLen(expectedContainersBefore))

		By("triggering the build again")
		buildSession = Fly.Start("trigger-job", "-w", "-j", "pipe/get-10m")
		<-buildSession.Exited
		Expect(buildSession.ExitCode()).To(Equal(1))

		By("expecting additional check containers for the task's image check and nested resource type check.")
		Expect(ContainersBy("type", "check")).To(HaveLen(6))

		By("expecting to only have new containers for build task image check, nested resource type check and build task")
		Expect(FlyTable("containers")).Should(HaveLen(expectedContainersBefore + 3))
	})
})

var _ = Describe("Tagged resource types", func() {
	BeforeEach(func() {
		Deploy("deployments/concourse.yml", "-o", "operations/tagged-worker.yml")

		By("setting a pipeline with tagged custom types")
		Fly.Run("set-pipeline", "-n", "-c", "pipelines/tagged-custom-types.yml", "-p", "pipe")
		Fly.Run("unpause-pipeline", "-p", "pipe")
	})

	It("is able to be used with tagged workers", func() {
		By("running a check which uses the tagged custom resource")
		Eventually(Fly.Start("check-resource", "-r", "pipe/10m")).Should(gexec.Exit(0))

		By("triggering a build which uses the tagged custom resource")
		buildSession := Fly.Start("trigger-job", "-w", "-j", "pipe/get-10m")
		<-buildSession.Exited
		Expect(buildSession.ExitCode()).To(Equal(0))
	})
})
