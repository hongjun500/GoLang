package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func main() {
	// 创建应用和窗口
	myApp := app.New()
	myWindow := myApp.NewWindow("Material Design Todo App")
	myWindow.Resize(fyne.NewSize(400, 600))

	// 数据：任务列表
	var tasks []string
	taskList := widget.NewList(
		func() int { return len(tasks) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Task Item")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(tasks[id])
		},
	)

	// 输入框和添加按钮
	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("Add a new task...")
	//addButton := material.NewButton(myWindow, "Add", func() { ... })

	addButton := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		taskText := strings.TrimSpace(taskEntry.Text)
		if taskText != "" {
			tasks = append(tasks, taskText)
			taskList.Refresh()
			taskEntry.SetText("")
		}
	})

	// 删除按钮（Material Design 风格的图标按钮）
	deleteButton := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		if len(tasks) == 0 {
			dialog.ShowInformation("Info", "No tasks to delete", myWindow)
			return
		}
		/*dialog.ShowConfirm("Delete", "Delete the selected task?", func(ok bool) {
			if ok && taskList.SelectedRow() >= 0 {
				index := taskList.SelectedRow()
				tasks = append(tasks[:index], tasks[index+1:]...)
				taskList.Refresh()
			}
		}, myWindow)*/
	})

	// 顶部工具栏（Material Design 的 AppBar）
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			dialog.ShowInformation("About", "Material Design Todo App\nPowered by Fyne", myWindow)
		}),
	)

	// 布局
	content := container.NewBorder(
		toolbar, // 顶部：工具栏
		container.NewVBox( // 底部：输入框和按钮
			taskEntry,
			container.NewHBox(addButton, deleteButton),
		),
		nil, nil,
		taskList, // 中间：任务列表
	)

	// 设置窗口内容并运行
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
