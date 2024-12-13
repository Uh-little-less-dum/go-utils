package signals

import (
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

type SetStageMsg struct {
	err      error
	NewStage build_stages.BuildStage
}

func SetStage(newStage build_stages.BuildStage) tea.Cmd {
	return func() tea.Msg {
		return SetStageMsg{
			err:      nil,
			NewStage: newStage,
		}
	}
}

type SetUseSelectedDirMsg struct {
	err            error
	UseSelectedDir bool
}

func SetUseSelectedDir(shouldUse bool) tea.Cmd {
	return func() tea.Msg {
		return SetUseSelectedDirMsg{
			err:            nil,
			UseSelectedDir: shouldUse,
		}
	}
}

type SetAcceptedTargetDirMsg struct {
	err       error
	TargetDir string
}

func SetAcceptedTargetDir(dirPath string) tea.Cmd {
	return func() tea.Msg {
		return SetAcceptedTargetDirMsg{
			err:       nil,
			TargetDir: dirPath,
		}
	}
}

type SetQuittingMsg struct {
	err error
}

func SetQuittingMessage(err error) tea.Cmd {
	return func() tea.Msg {
		return SetQuittingMsg{
			err: err,
		}
	}
}

type StdOutWrapperOutputMsg struct {
	Body string
}

func SendStdOutWrapperOutputMsg(content string) tea.Cmd {
	return func() tea.Msg {
		return StdOutWrapperOutputMsg{
			Body: content,
		}
	}
}

type SetConfigLocMethod_WaitForClone struct {
}

func SendConfigLocMethod_WaitForClone() tea.Cmd {
	return func() tea.Msg {
		return SetConfigLocMethod_WaitForClone{}
	}
}

type SetConfigLocMethod_pickFile struct {
}

func SendConfigLocMethod_pickFile() tea.Cmd {
	return func() tea.Msg {
		return SetConfigLocMethod_pickFile{}
	}
}

func SendUseEnvConfigResponse(wasAccepted bool) tea.Cmd {
	return func() tea.Msg {
		if wasAccepted {
			return SetStageMsg{
				err:      nil,
				NewStage: build_stages.CloneTemplateAppStage,
			}
		} else {
			return SetStageMsg{
				err:      nil,
				NewStage: build_stages.ChooseWaitOrPickConfigLoc,
			}
		}
	}
}

type RunSubCommandStreamMsg struct {
	StreamId stream_ids.StreamId
}

func SendRunSubCommandStreamMsg(streamId stream_ids.StreamId) tea.Cmd {
	return func() tea.Msg {
		return RunSubCommandStreamMsg{
			StreamId: streamId,
		}
	}
}

type SubCommandCompleteMsg struct {
	StreamId stream_ids.StreamId
	SubCmdId sub_command_ids.SubCommandId
}

func SendSubCommandCompleteMsg(streamId stream_ids.StreamId, subCommandId sub_command_ids.SubCommandId) tea.Cmd {
	return func() tea.Msg {
		return SubCommandCompleteMsg{
			StreamId: streamId,
			SubCmdId: subCommandId,
		}
	}
}

type FinishInitialTemplateCloneMsg struct {
}

func SendFinishInitialTemplateCloneMsg() tea.Cmd {
	return func() tea.Msg {
		return FinishInitialTemplateCloneMsg{}
	}
}
