package models

import (
  "bufio"
  "errors"
  "regexp"
  "strconv"
  "bytes"
  "strings"

  "github.com/youngjinpark20/gosu/modules/app"

  "github.com/teris-io/shortid"
)

type OsuMap struct {
  MapData
  General General
  Editor Editor
  Metadata Metadata
  Difficulty Difficulty
  TimingPoints []TimingPoint
  HitObjects []HitObject
}

type General struct {
  AudioFilename string
  AudioLeadIn int
  PreviewTime int
  Countdown int
  SampleSet string
  StackLeniency float64
  Mode int
  LetterboxInBreaks int
  StoryFireInFront int
  SkinPreference string
  EpilepsyWarning int
  CountdownOffset int
  WidescreenStoryboard int
  SpecialStyle int
  UseSkinSprites int
}

type Editor struct {
  Bookmarks []int
  DistanceSpacing float64
  BeatDivisor int
  GridSize int
  TimelineZoom float64
}

type Metadata struct {
  Title string
  TitleUnicode string
  Artist string
  ArtistUnicode string
  Creator string
  Version string
  Source string
  Tags []string
  BeatmapID int
  BeatmapSetID int
}

type Difficulty struct {
  HPDrainRate float64
  CircleSize float64
  OverallDifficulty float64
  ApproachRate float64
  SliderMultiplier float64
  SliderTickRate float64
}

type TimingPoint struct {
  Offset int
  MilliPerBeat float64
  Meter int
  SampleSet int
  SampleIndex int
  Volume int
  Inherited int
  KiaiMode int
}

type HitObject struct {
  X int
  Y int
  Time int
  Type int
  HitSounds int
  Extras string
}

func (m *OsuMap) Init() error {
  var err error
  m.ShortID, err = shortid.Generate()
  return err
}

func (ms *MapSet) ParseOsuMap(ctx *app.Context, data *bytes.Buffer) error {
  var _m Map
  var m OsuMap
  var i int
  var err error
  var section string

  _m.Init()
  m.Init()

  scanner := bufio.NewScanner(data)
  sectionTester := regexp.MustCompile("^\\[(.*)\\]$")
  keyValueTester := regexp.MustCompile("^(.*):\\s*(.*)\\s*$")
  timingPointTester := regexp.MustCompile("^(\\d*),([-|\\d]*),(\\d*),(\\d*),(\\d*),(\\d*),(\\d*),(\\d*)$")
  // hitObjectTester := regexp.MustCompile("^(\\d*),(\\d*),(\\d*),(\\d*),(\\d*),(.*)$")

  for scanner.Scan() {
    line := scanner.Text()
    if i<200 {
      i++
    }

    if sectionTester.MatchString(line) {
      section = sectionTester.FindStringSubmatch(line)[1]
    } else {
      switch section {
      case "General", "Editor", "Metadata", "Difficulty":
        if keyValueTester.MatchString(line) {
          submatches := keyValueTester.FindStringSubmatch(line);
          key := submatches[1]
          value := submatches[2]

          switch key {
          // GENERAL
          case "AudioFilename":         m.General.AudioFilename = value
          case "AudioLeadIn":           m.General.AudioLeadIn, err = strconv.Atoi(value)
          case "PreviewTime":           m.General.PreviewTime, err = strconv.Atoi(value)
          case "Countdown":             m.General.Countdown, err = strconv.Atoi(value)
          case "SampleSet":             m.General.SampleSet = value
          case "StackLeniency":         m.General.StackLeniency, err = strconv.ParseFloat(value, 64)
          case "Mode":                  m.General.Mode, err = strconv.Atoi(value)
          case "LetterboxInBreaks":     m.General.LetterboxInBreaks, err = strconv.Atoi(value)
          case "StoryFireInFront":      m.General.StoryFireInFront, err = strconv.Atoi(value)
          case "SkinPreference":        m.General.SkinPreference = value
          case "EpilepsyWarning":       m.General.EpilepsyWarning, err = strconv.Atoi(value)
          case "CountdownOffset":       m.General.CountdownOffset, err = strconv.Atoi(value)
          case "WidescreenStoryboard":  m.General.WidescreenStoryboard, err = strconv.Atoi(value)
          case "SpecialStyle":          m.General.SpecialStyle, err = strconv.Atoi(value)
          case "UseSkinSprites":        m.General.UseSkinSprites, err = strconv.Atoi(value)
          // EDITOR
          case "Bookmarks":
            bookmarkStrings := strings.Split(value, ",")
            for _, elem := range bookmarkStrings {
              var bookmark int
              bookmark, err = strconv.Atoi(elem)
              m.Editor.Bookmarks = append(m.Editor.Bookmarks, bookmark)
              if err != nil {
                return err
              }
            }
          case "DistanceSpacing":  m.Editor.DistanceSpacing, err = strconv.ParseFloat(value, 64)
          case "BeatDivisor":      m.Editor.BeatDivisor, err = strconv.Atoi(value)
          case "GridSize":         m.Editor.GridSize, err = strconv.Atoi(value)
          case "TimelineZoom":     m.Editor.TimelineZoom, err = strconv.ParseFloat(value, 64)
          // METADATA
          case "Title":         m.Metadata.Title = value
          case "TitleUnicode":  m.Metadata.TitleUnicode = value
          case "Artist":        m.Metadata.Artist = value
          case "ArtistUnicode": m.Metadata.ArtistUnicode = value
          case "Creator":       m.Metadata.Creator = value
          case "Version":       m.Metadata.Version = value
          case "Source":        m.Metadata.Source = value
          case "Tags":          m.Metadata.Tags = strings.Split(value, " ")
          case "BeatmapID":     m.Metadata.BeatmapID, err = strconv.Atoi(value)
          case "BeatmapSetID":  m.Metadata.BeatmapSetID, err = strconv.Atoi(value)
          // DIFFICULTY
          case "HPDrainRate":       m.Difficulty.HPDrainRate, err = strconv.ParseFloat(value, 64)
          case "CircleSize":        m.Difficulty.CircleSize, err = strconv.ParseFloat(value, 64)
          case "OverallDifficulty": m.Difficulty.OverallDifficulty, err = strconv.ParseFloat(value, 64)
          case "ApproachRate":      m.Difficulty.ApproachRate, err = strconv.ParseFloat(value, 64)
          case "SliderMultiplier":  m.Difficulty.SliderMultiplier, err = strconv.ParseFloat(value, 64)
          case "SliderTickRate":    m.Difficulty.SliderTickRate, err = strconv.ParseFloat(value, 64)
          default: return errors.New("Unhandled key-value pair (" + key + ": " + value + ")")
          }
          if err != nil {
            return err
          }
        }

      case "Events": // TODO not implemented
      case "Colours": // TODO not implemented

      case "TimingPoints":
        if timingPointTester.MatchString(line) {
          var tp TimingPoint
          submatches := timingPointTester.FindStringSubmatch(line)
          tp.Offset, err = strconv.Atoi(submatches[1])
          tp.MilliPerBeat, err = strconv.ParseFloat(submatches[2], 64)
          tp.Meter, err = strconv.Atoi(submatches[3])
          tp.SampleSet, err = strconv.Atoi(submatches[4])
          tp.SampleIndex, err = strconv.Atoi(submatches[5])
          tp.Volume, err = strconv.Atoi(submatches[6])
          tp.Inherited, err = strconv.Atoi(submatches[7])
          tp.KiaiMode, err = strconv.Atoi(submatches[8])
          m.TimingPoints = append(m.TimingPoints, tp)
        }

      case "HitObjects":
        if timingPointTester.MatchString(line) {
          var tp TimingPoint
          submatches := timingPointTester.FindStringSubmatch(line)
          tp.Offset, err = strconv.Atoi(submatches[1])
          tp.MilliPerBeat, err = strconv.ParseFloat(submatches[2], 64)
          tp.Meter, err = strconv.Atoi(submatches[3])
          tp.SampleSet, err = strconv.Atoi(submatches[4])
          tp.SampleIndex, err = strconv.Atoi(submatches[5])
          tp.Volume, err = strconv.Atoi(submatches[6])
          tp.Inherited, err = strconv.Atoi(submatches[7])
          tp.KiaiMode, err = strconv.Atoi(submatches[8])
          m.TimingPoints = append(m.TimingPoints, tp)
        }

      case "": // Do nothing
      default: return errors.New("Unmatched section " + section)
      }
    }
  }

  _m = Map{
    ShortID:     _m.ShortID,
    Title:       m.Metadata.Title,
    TitleOther:  m.Metadata.TitleUnicode,
    Artist:      m.Metadata.Artist,
    ArtistOther: m.Metadata.ArtistUnicode,
    Creator:     m.Metadata.Creator,
    Version:     m.Metadata.Version,
    Tags:        m.Metadata.Tags,
    Type:        "osu",
    ParentID:    ms.ShortID,
    DataID:      m.ShortID,
  }

  ms.Title =       m.Metadata.Title
  ms.TitleOther =  m.Metadata.TitleUnicode
  ms.Artist =      m.Metadata.Artist
  ms.ArtistOther = m.Metadata.ArtistUnicode
  ms.Creator =     m.Metadata.Creator
  ms.Tags =        m.Metadata.Tags

  ctx.DB.Create(&_m)
  ctx.DB.Create(&m)

  ms.MapIDs = append(ms.MapIDs, _m.ShortID)

  return nil
}

