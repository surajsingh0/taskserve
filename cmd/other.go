package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todo/internal/storage_type"
	"todo/internal/task"

	"github.com/spf13/cobra"
)

var tm *task.TaskManager

func Close() {
	if tm != nil {
		tm.Close()
	}
}

func init() {
	var err error
	tm, err = task.NewTaskManager(storage_type.CSV)
	if err != nil {
		fmt.Println("Error initializing task manager:", err)
		return
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(toggleCompletedCmd)
	rootCmd.AddCommand(totalCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(webServerCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run:   addCmdRun,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run:   listCmdRun,
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run:   deleteCmdRun,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task with a new task",
	Run:   updateCmdRun,
}

var toggleCompletedCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle task's completed status",
	Run:   toggleCmdRun,
}

var totalCmd = &cobra.Command{
	Use:   "total",
	Short: "Get the total number of tasks",
	Run:   totalCmdRun,
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear/Delete all the tasks",
	Run:   clearCmdRun,
}

var webServerCmd = &cobra.Command{
	Use:   "webserver",
	Short: "Start a web server to manage tasks through a web interface",
	Run:   webServerCmdRun,
}

func addCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide a task to add.")
		return
	}
	task := args[0]

	if err := tm.AddTask(task, false); err != nil {
		fmt.Println("Error adding task:", err)
	} else {
		fmt.Printf("Task '%s' added!\n", task)
	}
}

func listCmdRun(cmd *cobra.Command, args []string) {
	tasksList, err := tm.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks:", err)
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasksList {
		fmt.Printf("- ID: %d, Title: %s, Completed: %t, Date: %s\n",
			task.Id, task.Title, task.Completed, task.Date.Format("2006-01-02 15:04:05"))
	}
}

func totalCmdRun(cmd *cobra.Command, args []string) {
	count, err := tm.TotalTasks()
	if err != nil {
		fmt.Println("Error getting total tasks:", err)
		return
	}
	fmt.Printf("Total tasks: %d\n", count)
}

func deleteCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide a task ID to delete.")
		return
	}
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		return
	}
	if err := tm.DeleteTask(taskId); err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}
	fmt.Printf("Task deleted: %d\n", taskId)
}

func updateCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide an existing task ID and a new task.")
		return
	}
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		return
	}
	newTask := args[1]
	if err := tm.UpdateTask(taskId, newTask); err != nil {
		fmt.Println("Error updating task:", err)
		return
	}
	fmt.Printf("Task updated: %s\n", newTask)
}

func toggleCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide an existing task ID.")
		return
	}
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		return
	}
	if err := tm.ToggleCompleted(taskId); err != nil {
		fmt.Println("Error toggling task's status:", err)
		return
	}
	fmt.Printf("Task toggled: %d\n", taskId)
}

func clearCmdRun(cmd *cobra.Command, args []string) {
	err := tm.Clear()
	if err != nil {
		fmt.Println("Failed to clear the tasks:", err)
		return
	}
	fmt.Println("All the tasks got successfully cleared/deleted!")
}

func webServerCmdRun(cmd *cobra.Command, args []string) {
	startWebServer()
}
